package kiss

import (
	"fmt"
	"io"
)

const (
	fend  = 0xc0
	fesc  = 0xdb
	tfend = 0xdc
	tfesc = 0xdd
)

// Escape replace TNC special characters with non-control sequences.
func Escape(data []byte) []byte {
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case fend:
			data[i] = fesc
			data = append(data[:i+1], data[i:]...)
			data[i+1] = tfend
		case fesc:
			data[i] = fesc
			data = append(data[:i+1], data[i:]...)
			data[i+1] = tfesc
		}
	}
	return data
}

// Unescape replace escape sequences with original data bytes.
func Unescape(data []byte) ([]byte, error) {
	for i := 0; i < len(data); i++ {
		if data[i] == fesc {
			if i+1 == len(data) {
				return data, io.ErrUnexpectedEOF
			}

			switch data[i+1] {
			case tfend:
				data[i] = fend
			case tfesc:
				data[i] = fesc
			default:
				return data, fmt.Errorf("unknown escape sequence 0x%02x", data[i+1])
			}

			data = append(data[:i+1], data[i+2:]...)
		}
	}
	return data, nil
}
