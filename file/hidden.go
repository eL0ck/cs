// SPDX-License-Identifier: MIT
// SPDX-License-Identifier: Unlicense
// +build !windows

package file

import (
	"os"
)

// IsHidden Returns true if file is hidden
func IsHidden(file os.FileInfo, directory string) (bool, error) {
	return file.Name()[0:1] == ".", nil
}
