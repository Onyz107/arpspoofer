//go:build linux

package sysctl

var sysctlValues = map[string]int{
	"net.ipv4.conf.all.forwarding": 1,
	"net.ipv4.ip_forward":          1,
	"net.ipv6.conf.all.forwarding": 1,
	"net.ipv4.conf.all.rp_filter":  0,
}
