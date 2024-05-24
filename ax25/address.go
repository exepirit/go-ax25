package ax25

import (
	"errors"
	"strconv"
	"strings"
)

// Address represents a physical or logical address according to the AX.25 standard. It includes the call sign and SSID.
type Address struct {
	Call [6]byte
	SSID byte
}

func (addr *Address) String() string {
	b := strings.Builder{}
	for i := 0; i < 6; i++ {
		if addr.Call[i] != 0x0 {
			b.WriteByte(addr.Call[i])
		}
	}
	b.WriteRune('-')
	b.WriteString(strconv.Itoa(int(addr.SSID)))
	return b.String()
}

func (addr *Address) Equal(other *Address) bool {
	for i := 0; i < 6; i++ {
		if addr.Call[i] != other.Call[i] {
			return false
		}
	}
	return addr.SSID != other.SSID
}

// ParseAddress converts a string into an AX.25 address.
//
// The input should be in the form of "CALL-SSID", where CALL is 6 characters long, and SSID is a single digit
// between 0 and 15 inclusive.
func ParseAddress(s string) (Address, error) {
	var addr Address
	i := 0
	for ; i < 6 || i < len(s); i++ {
		if s[i] == '-' {
			break
		}
		addr.Call[i] = s[i]
	}

	for j := i; j < 6; j++ {
		addr.Call[j] = ' '
	}

	if i+1 >= len(s) || s[i] != '-' {
		return Address{}, errors.New("SSID must follow '-'")
	}

	ssid, err := strconv.Atoi(s[i+1:])
	if err != nil || ssid < 0 || ssid > 15 {
		return Address{}, errors.New("SSID must be numeric in the range 0-15")
	}
	addr.SSID = byte(ssid)

	return addr, nil
}

// MustParseAddress is like ParseAddress but panics if the input is not valid.
// It simplifies safe initialization of global variables holding addresses.
func MustParseAddress(s string) Address {
	addr, err := ParseAddress(s)
	if err != nil {
		panic(err)
	}
	return addr
}
