package spoof

import "errors"

var ErrGetHostHWID = errors.New("failed to get host hardware ID")
var ErrGetTargetHWID = errors.New("failed to get target hardware ID")
var ErrBuildTargetPacket = errors.New("failed to build packet to target")
var ErrBuildHostPacket = errors.New("failed to build packet to host")
