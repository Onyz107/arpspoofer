package handle

import "errors"

var ErrOpenInterface = errors.New("failed to open network interface")
var ErrGetHWID = errors.New("failed to get hardware ID of interface")
var ErrGetIPAddr = errors.New("failed to get IP address of interface")
var ErrHWIDNotFound = errors.New("hardware ID not found for interface")
var ErrIfaceName = errors.New("invalid interface name")
var ErrIfaceAddrs = errors.New("failed to get addresses for interface")
var ErrNoIPAddr = errors.New("no valid IPv4 address found for interface")
