//+build linux

package smbios

import (
	"io"
	"os"
)

const (
	// sysfs locations for SMBIOS information.
	sysfsDMI        = "/sys/firmware/dmi/tables/DMI"
	sysfsEntryPoint = "/sys/firmware/dmi/tables/smbios_entry_point"
)

// stream opens the SMBIOS entry point and an SMBIOS structure stream.
func stream() (io.ReadCloser, EntryPoint, error) {
	// First, check for the sysfs location present in modern kernels.
	_, err := os.Stat(sysfsEntryPoint)
	switch {
	case err == nil:
		return sysfsStream(sysfsEntryPoint, sysfsDMI)
	case os.IsNotExist(err):
		// Fall back to the standard UNIX-like system method.
		return devMemStream()
	default:
		return nil, nil, err
	}
}

// sysfsStream reads the SMBIOS entry point and structure stream from
// two files; usually the modern sysfs locations.
func sysfsStream(entryPoint, dmi string) (io.ReadCloser, EntryPoint, error) {
	epf, err := os.Open(entryPoint)
	if err != nil {
		return nil, nil, err
	}
	defer epf.Close()

	ep, err := ParseEntryPoint(epf)
	if err != nil {
		return nil, nil, err
	}

	sf, err := os.Open(dmi)
	if err != nil {
		return nil, nil, err
	}

	return sf, ep, nil
}
