package sysctl

import "errors"

var ErrReadFile = errors.New("failed to read file")
var ErrConvertStringToInt = errors.New("failed to convert string to int")
var ErrGetSysctl = errors.New("failed to get sysctl value")
var ErrInvalidValue = errors.New("invalid value")
var ErrSyscallSysctl = errors.New("syscall sysctl error")
