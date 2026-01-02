//go:build linux || darwin

package sysctl

import (
	"errors"
	"fmt"
)

func CheckSysctl() error {
	var errs []error

	for key, expectedValue := range sysctlValues {
		value, err := getSysctlValue(key)
		if err != nil {
			errs = append(errs, errors.Join(ErrGetSysctl, fmt.Errorf("%s", key), err))
			continue
		}
		if value != expectedValue {
			errs = append(errs, errors.Join(ErrInvalidValue, fmt.Errorf("%s: expected %d, got %d", key, expectedValue, value)))
			continue
		}
	}
	return errors.Join(errs...)
}
