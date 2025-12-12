package hwid_test

import (
	"net"
	"testing"

	"github.com/Onyz107/arpspoofer/handle"
	"github.com/Onyz107/arpspoofer/internal/hwid"
)

func TestGetFromIP(t *testing.T) {
	ifaceHandle, err := handle.Open("wlx00c0cab2706d")
	if err != nil {
		t.Fatal(err)
	}
	defer ifaceHandle.Close()

	// Construct gatway IP
	gatewayIP := make(net.IP, len(ifaceHandle.IP))
	copy(gatewayIP, ifaceHandle.IP)
	gatewayIP[15] = 1
	t.Logf("Constructed gatway IP as: %s", gatewayIP.To4().String())

	mac, err := hwid.GetFromIP(t.Context(), ifaceHandle, gatewayIP, true)
	if err != nil {
		t.Fatalf("Failed to get HWID: %v", err)
	}

	t.Logf("Gateway MAC Address: %s", mac.String())
}
