package handle

import "errors"

var (
	ErrOpenInterface = errors.New("failed to open network interface")
	ErrGetHWID       = errors.New("failed to get hardware ID of interface")
	ErrGetIPAddr     = errors.New("failed to get IP address of interface")
	ErrHWIDNotFound  = errors.New("hardware ID not found for interface")
	ErrIfaceName     = errors.New("invalid interface name")
	ErrIfaceAddrs    = errors.New("failed to get addresses for interface")
	ErrNoIPAddr      = errors.New("no valid IPv4 address found for interface")
)
