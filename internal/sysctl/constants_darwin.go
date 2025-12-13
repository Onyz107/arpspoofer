//go:build darwin

package sysctl

var sysctlValues = map[string]int{
	"net.inet.ip.forwarding":   1,
	"net.inet6.ip6.forwarding": 1,
}
