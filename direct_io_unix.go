// Direct IO for Unix

// +build !windows,!darwin,!openbsd,!plan9

package directio

import (
	"os"
	"syscall"
)

const (
	// Size to align the buffer to
	AlignSize = 512

	// Minimum block size
	BlockSize = 16 * 1024
)

// OpenFile is a modified version of os.OpenFile which sets O_DIRECT
func OpenFile(name string, flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(name, syscall.O_DIRECT|flag, perm)
}
