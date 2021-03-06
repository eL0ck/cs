// SPDX-License-Identifier: MIT
// SPDX-License-Identifier: Unlicense
// +build windows

package file

import (
	"os"
	"path"
	"syscall"
)

// IsHidden Returns true if file is hidden
func IsHidden(file os.FileInfo, directory string) (bool, error) {
	fullpath := path.Join(directory, file.Name())
	pointer, err := syscall.UTF16PtrFromString(fullpath)
	if err != nil {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
