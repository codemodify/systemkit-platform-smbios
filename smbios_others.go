//+build !dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows

package smbios

import (
	"fmt"
	"io"
	"runtime"
)

// stream is not implemented for unsupported platforms.
func stream() (io.ReadCloser, EntryPoint, error) {
	return nil, nil, fmt.Errorf("opening SMBIOS stream not implemented on %q", runtime.GOOS)
}
