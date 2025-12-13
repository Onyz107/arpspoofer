//go:build darwin

package sysctl

import (
	"errors"
	"syscall"
)

func getSysctlValue(key string) (int, error) {
	val, err := syscall.SysctlUint32(key)
	if err != nil {
		return 0, errors.Join(ErrSyscallSysctl, err)
	}
	return int(val), nil
}
