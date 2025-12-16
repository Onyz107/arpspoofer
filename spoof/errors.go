package spoof

import "errors"

var (
	ErrGetHostHWID       = errors.New("failed to get host hardware ID")
	ErrGetTargetHWID     = errors.New("failed to get target hardware ID")
	ErrBuildTargetPacket = errors.New("failed to build packet to target")
	ErrBuildHostPacket   = errors.New("failed to build packet to host")
)
