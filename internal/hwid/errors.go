package hwid

import "errors"

var ErrBuildPkt = errors.New("failed to build ARP packet")
var ErrWritePacketData = errors.New("failed to write packet data to interface")
var ErrFindHWID = errors.New("failed to find hardware ID for given IP")
var ErrTimeout = errors.New("timeout waiting for ARP reply")
