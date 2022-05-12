package tnc_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/exepirit/go-ax25/pkg/tnc"
)

func TestEscape(t *testing.T) {
	testCases := []struct {
		Data    []byte
		Escaped []byte
	}{
		{
			Data:    []byte("TEST"),
			Escaped: []byte{0x54, 0x45, 0x53, 0x54},
		},
		{
			Data:    []byte("Hello"),
			Escaped: []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f},
		},
		{
			Data:    []byte{0xc0, 0xdb},
			Escaped: []byte{0xdb, 0xdc, 0xdb, 0xdd},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := tnc.Escape(testCase.Data)

			if len(result) != len(testCase.Escaped) {
				t.Fatalf("invalid array size. expected: %d, got: %d", len(testCase.Escaped), len(result))
				return
			}

			for i := 0; i < len(result); i++ {
				if result[i] != testCase.Escaped[i] {
					t.Logf("expected: %v\ngot: %v", testCase.Escaped, result)
					t.Fatalf("expected: 0x%02x, got: 0x%02x on position %d", testCase.Escaped[i], result[i], i)
					return
				}
			}
		})
	}
}

func TestUnescape(t *testing.T) {
	testCases := []struct {
		Escaped       []byte
		Unescaped     []byte
		ExpectedError error
	}{
		{
			Escaped:       []byte{0xdb, 0xdc, 0xdb, 0xdd},
			Unescaped:     []byte{0xc0, 0xdb},
			ExpectedError: nil,
		},
		{
			Escaped:       []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f},
			Unescaped:     []byte("Hello"),
			ExpectedError: nil,
		},
		{
			Escaped:       []byte{0x00, 0xdb},
			ExpectedError: io.ErrUnexpectedEOF,
		},
		{
			Escaped:       []byte{0x00, 0x0db, 0xb1, 0x02},
			ExpectedError: errors.New("unknown escape sequence 0xb1"),
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result, err := tnc.Unescape(testCase.Escaped)

			if err != nil {
				if testCase.ExpectedError != nil {
					if err.Error() != testCase.ExpectedError.Error() {
						t.Fatalf("expected error: %s, got: %s", testCase.ExpectedError, err)
					}
				} else {
					t.Fatalf("unexpected error: %s", err)
				}
				return
			}

			if len(result) != len(testCase.Unescaped) {
				t.Fatalf("invalid array size. expected: %d, got: %d", len(testCase.Unescaped), len(result))
				return
			}

			for i := 0; i < len(result); i++ {
				if result[i] != testCase.Unescaped[i] {
					t.Logf("expected: %v\ngot: %v", testCase.Unescaped, result)
					t.Fatalf("expected: 0x%02x, got: 0x%02x on position %d", testCase.Unescaped[i], result[i], i)
					return
				}
			}
		})
	}
}
