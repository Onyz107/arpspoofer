package sysctl

import "errors"

var (
	ErrReadFile           = errors.New("failed to read file")
	ErrConvertStringToInt = errors.New("failed to convert string to int")
	ErrGetSysctl          = errors.New("failed to get sysctl value")
	ErrInvalidValue       = errors.New("invalid value")
	ErrSyscallSysctl      = errors.New("syscall sysctl error")
)
