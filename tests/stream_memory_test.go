package tests

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"

	smbios "github.com/codemodify/systemkit-platform-smbios"
	"github.com/google/go-cmp/cmp"
)

func Test_MemoryStream(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		ss   []*smbios.Structure
		ok   bool
	}{
		{
			name: "empty",
			b:    nil,
		},
		{
			name: "magic before first paragraph",
			b: makeMemory(
				[]byte{'_', 'S', 'M', '_'},
				nil,
				nil,
			),
		},
		{
			name: "magic after last paragraph",
			b: makeMemory(
				nil,
				nil,
				[]byte{'_', 'S', 'M', '_'},
			),
		},
		{
			name: "64, OK",
			b: func() []byte {
				// Just enough information to point to an address
				// that contains the structure stream.
				const addr = 0x00f0
				epb := mustMarshalEntryPoint(&smbios.EntryPoint64Bit{
					StructureTableMaxSize: 512,
					StructureTableAddress: addr,
				})

				// Place entry point in searchable range.
				b := makeMemory(
					nil,
					epb,
					nil,
				)

				// Structure stream, placed starting at the address
				// specified in entry point.
				stream := []byte{
					0x00, 0x05, 0x01, 0x00,
					0xff,
					0x00,
					0x00,

					0x01, 0x0c, 0x02, 0x00,
					0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
					'd', 'e', 'a', 'd', 'b', 'e', 'e', 'f', 0x00,
					0x00,

					127, 0x06, 0x03, 0x00,
					0x01, 0x02,
					'a', 'b', 'c', 'd', 0x00,
					'1', '2', '3', '4', 0x00,
					0x00,
				}

				copy(b[addr:], stream)

				return b
			}(),
			ss: []*smbios.Structure{
				{
					Header: smbios.Header{
						Type:   0,
						Length: 5,
						Handle: 1,
					},
					Formatted: []byte{0xff},
				},
				{
					Header: smbios.Header{
						Type:   1,
						Length: 12,
						Handle: 2,
					},
					Formatted: []byte{0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef},
					Strings:   []string{"deadbeef"},
				},
				{
					Header: smbios.Header{
						Type:   127,
						Length: 6,
						Handle: 3,
					},
					Formatted: []byte{0x01, 0x02},
					Strings:   []string{"abcd", "1234"},
				},
			},
			ok: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := bytes.NewReader(tt.b)

			rc, _, err := smbios.MemoryStream(rs, start, end)

			if tt.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tt.ok && err == nil {
				t.Fatalf("expected an error, but none occurred: %v", err)
			}

			if !tt.ok {
				// Don't bother doing comparison if entry point is invalid.
				t.Logf("OK error: %v", err)
				return
			}
			defer rc.Close()

			ss, err := smbios.NewDecoder(rc).Decode()
			if err != nil {
				t.Fatalf("failed to decode structures: %v", err)
			}

			if diff := cmp.Diff(tt.ss, ss); diff != "" {
				t.Fatalf("unexpected structures (-want +got):\n%s", diff)
			}
		})
	}
}

// Memory addresses used to start and stop searching for entry points.
const (
	start = 0x0010
	end   = 0xfff0
)

func makeMemory(before, in, after []byte) []byte {
	b := make([]byte, math.MaxUint16)

	copy(b[0x0000:start], before)
	copy(b[start:0xfff0], in)
	copy(b[end:0xffff], after)

	return b
}

func mustMarshalEntryPoint(ep smbios.EntryPoint) []byte {
	switch x := ep.(type) {
	case *smbios.EntryPoint64Bit:
		return marshal64(x)
	default:
		// TODO(mdlayher): expand with 32-bit entry point.
		panic(fmt.Sprintf("entry point marshaling not implemented for %T", ep))
	}
}

func marshal64(ep *smbios.EntryPoint64Bit) []byte {
	b := make([]byte, smbios.ExpLen64)

	copy(b[0:5], smbios.Magic64)
	b[6] = smbios.ExpLen64

	b[7] = ep.Major
	b[8] = ep.Minor
	b[9] = ep.Revision
	b[10] = ep.EntryPointRevision
	b[11] = ep.Reserved
	binary.LittleEndian.PutUint32(b[12:16], ep.StructureTableMaxSize)
	binary.LittleEndian.PutUint64(b[16:24], ep.StructureTableAddress)

	var chk uint8
	for i := range b {
		// Explicitly skip the checksum byte for computation.
		if i == smbios.ChkIndex64 {
			continue
		}

		chk += b[i]
	}

	// Produce the correct checksum for the entry point.
	b[5] = uint8(256 - int(chk))

	return b
}
