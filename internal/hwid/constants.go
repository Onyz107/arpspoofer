package hwid

import (
	"net"
	"time"
)

const tickerTime = 350 * time.Millisecond
const timeoutTime = 5 * time.Second

var broadcastMAC = net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
