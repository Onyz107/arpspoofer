package hwid

import "errors"

var (
	ErrBuildPkt        = errors.New("failed to build ARP packet")
	ErrWritePacketData = errors.New("failed to write packet data to interface")
	ErrFindHWID        = errors.New("failed to find hardware ID for given IP")
	ErrTimeout         = errors.New("timeout waiting for ARP reply")
)
