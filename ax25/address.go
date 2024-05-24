package ax25

import (
	"errors"
	"strconv"
	"strings"
)

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

func MustParseAddress(s string) Address {
	addr, err := ParseAddress(s)
	if err != nil {
		panic(err)
	}
	return addr
}
