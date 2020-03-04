//+build dragonfly freebsd netbsd openbsd solaris

// Linux intentionally omitted because it has an alternative method that
// is used before attempting /dev/mem access.  See stream_linux.go.

package smbios

import (
	"io"
)

// stream opens the SMBIOS entry point and an SMBIOS structure stream.
func stream() (io.ReadCloser, EntryPoint, error) {
	// Use the standard UNIX-like system method.
	return devMemStream()
}
