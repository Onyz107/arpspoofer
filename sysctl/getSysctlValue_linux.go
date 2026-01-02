//go:build linux

package sysctl

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSysctlValue(key string) (int, error) {
	// Replace dots with slashes for procfs path
	path := fmt.Sprintf("/proc/sys/%s", strings.ReplaceAll(key, ".", "/"))

	output, err := os.ReadFile(path)
	if err != nil {
		return 0, errors.Join(ErrReadFile, err)
	}

	valueStr := strings.TrimSpace(string(output))

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, errors.Join(ErrConvertStringToInt, err)
	}

	return value, nil
}
