package main

import (
	"fmt"
	"log"

	smbios "github.com/codemodify/systemkit-platform-smbios"
)

// accesses and displays SMBIOS data.

func main() {
	// Find SMBIOS data in operating system-specific location.
	rc, ep, err := smbios.Stream()
	if err != nil {
		log.Fatalf("failed to open stream: %v", err)
	}
	// Be sure to close the stream!
	defer rc.Close()

	// Decode SMBIOS structures from the stream.
	d := smbios.NewDecoder(rc)
	ss, err := d.Decode()
	if err != nil {
		log.Fatalf("failed to decode structures: %v", err)
	}

	// Determine SMBIOS version and table location from entry point.
	major, minor, rev := ep.Version()
	addr, size := ep.Table()

	fmt.Printf("SMBIOS %d.%d.%d - table: address: %#x, size: %d\n", major, minor, rev, addr, size)

	for _, s := range ss {
		fmt.Println("###################################################")
		fmt.Println(fmt.Sprintf("Header.Type %v", s.Header.Type))
		fmt.Println(fmt.Sprintf("Header.Length %v", s.Header.Length))
		fmt.Println(fmt.Sprintf("Header.Handle %v", s.Header.Handle))

		// fmt.Println(fmt.Sprintf("Formatted %s", string(s.Formatted)))

		for i, st := range s.Strings {
			fmt.Println(fmt.Sprintf("Strings[%d] -> %s", i, st))
		}
	}
}
