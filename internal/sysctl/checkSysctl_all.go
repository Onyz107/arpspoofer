//go:build !linux && !darwin

package sysctl

// Added for compatibility; no sysctl checks on other OSes.

func CheckSysctl() error {
	return nil
}
