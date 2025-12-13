package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Onyz107/arpspoofer/handle"
	"github.com/Onyz107/arpspoofer/internal/sysctl"
	"github.com/Onyz107/arpspoofer/spoof"
	"github.com/urfave/cli/v2"
)

type options struct {
	HostIP    string
	TargetIP  string
	Interface string
	Interval  time.Duration
	Restore   bool
	Verbose   bool
}

func main() {
	opts := &options{}

	app := &cli.App{
		Name:    os.Args[0],
		Usage:   "A reliable ARP spoofer",
		Version: "v1.0.5",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "target",
				Aliases:     []string{"t"},
				Usage:       "IP address of the target to spoof",
				Required:    true,
				Destination: &opts.TargetIP,
			},
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"g"},
				Usage:       "IP address of the host/gateway to spoof",
				Required:    true,
				Destination: &opts.HostIP,
			},
			&cli.StringFlag{
				Name:        "interface",
				Aliases:     []string{"i"},
				Usage:       "Network interface to use",
				Required:    true,
				Destination: &opts.Interface,
			},
			&cli.DurationFlag{
				Name:        "interval",
				Aliases:     []string{"n"},
				Usage:       "Interval between ARP packets",
				Value:       time.Second,
				Destination: &opts.Interval,
			},
			&cli.BoolFlag{
				Name:        "restore",
				Aliases:     []string{"r"},
				Usage:       "Restore ARP tables when exiting",
				Value:       true,
				Destination: &opts.Restore,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Usage:       "Enable verbose output",
				Value:       false,
				Destination: &opts.Verbose,
			},
		},
		Action: func(c *cli.Context) error {
			if !isValidIPv4(opts.HostIP) {
				return ErrInvalidHostIP
			}
			if !isValidIPv4(opts.TargetIP) {
				return ErrInvalidTargetIP
			}
			if !isValidInterface(opts.Interface) {
				return ErrInvalidInterface
			}

			if opts.Verbose {
				log.Printf("Flags set:\n\t%#v\n", opts)
			}

			if err := sysctl.CheckSysctl(); err != nil {
				return errors.Join(ErrInvalidSysctlSettings, err)
			}

			ifaceHandle, err := handle.Open(opts.Interface)
			if err != nil {
				return errors.Join(ErrOpenInterface, err)
			}
			defer ifaceHandle.Close()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			stop, err := spoof.Spoof(ctx, ifaceHandle, net.ParseIP(opts.HostIP), net.ParseIP(opts.TargetIP), opts.Interval, opts.Restore, opts.Verbose)
			if err != nil {
				return errors.Join(ErrStartSpoofing, err)
			}

			log.Println("ARP spoofing started. Press Ctrl+C to stop.")
			sigCtx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
			defer cancel()

			<-sigCtx.Done()
			log.Println("Stopping ARP spoofing.")
			stop()
			log.Println("ARP spoofing stopped.")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func isValidIPv4(ip string) bool {
	ip4 := net.ParseIP(ip)
	return !(ip4 == nil || ip4.To4() == nil)
}

func isValidInterface(name string) bool {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return false
	}
	return iface.Flags&net.FlagUp != 0
}
