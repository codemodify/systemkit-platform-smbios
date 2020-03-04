//+build gofuzz

package tests

import (
	"bytes"
)

func Fuzz(data []byte) int {
	return fuzzDecoder(data)
}

func fuzzDecoder(data []byte) int {
	d := NewDecoder(bytes.NewReader(data))

	if _, err := d.Decode(); err != nil {
		return 0
	}

	return 1
}
