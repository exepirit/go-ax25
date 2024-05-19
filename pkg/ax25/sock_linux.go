//go:build cgo && linux && !nokernel
// +build cgo,linux,!nokernel

package ax25

import "C"
import (
	"fmt"
	"syscall"
)

// #include <sys/socket.h>
// #include <linux/ax25.h>
// #include <unistd.h>
//
// typedef struct sockaddr_ax25 sockaddr_ax25;
// typedef struct sockaddr sockaddr;
import "C"

import (
	"unsafe"
)

type socket struct {
	fd int
}

func (s *socket) recvFrom(data []byte) (n int, addr *Address, err error) {
	var dest C.sockaddr_ax25
	var destLen C.uint
	readed, err := C.recvfrom(
		C.int(s.fd),
		C.CBytes(data), C.size_t(len(data)),
		0,
		(*C.sockaddr)(unsafe.Pointer(&dest)), (*C.uint)(unsafe.Pointer(&destLen)),
	)
	if err != nil {
		return int(readed), nil, err.(syscall.Errno)
	}

	addr = &Address{}
	for i := 0; i < 6; i++ {
		addr.Call[i] = byte(dest.sax25_call.ax25_call[i])
	}
	addr.SSID = byte(dest.sax25_call.ax25_call[6])
	return int(readed), addr, nil
}

func (s *socket) sendTo(data []byte, addr *Address) (n int, err error) {
	dest := addressToSockaddr(addr)
	written, err := C.sendto(
		C.int(s.fd),
		C.CBytes(data), C.ulong(len(data)),
		0,
		(*C.sockaddr)(unsafe.Pointer(&dest)), C.uint(unsafe.Sizeof(dest)),
	)
	if err != nil {
		return int(written), err.(syscall.Errno)
	}
	return int(written), nil
}

func (s *socket) close() error {
	_, err := C.close(C.int(s.fd))
	if err != nil {
		return err.(syscall.Errno)
	}
	return nil
}

func openUnnumberedSocket(source *Address) (*socket, error) {
	fd, err := C.socket(C.AF_AX25, C.SOCK_DGRAM, 0)
	if err != nil {
		C.close(fd)
		return nil, fmt.Errorf("failed to open socket: %w", err.(syscall.Errno))
	}

	sockaddr := addressToSockaddr(source)
	_, err = C.bind(fd, (*C.sockaddr)(unsafe.Pointer(&sockaddr)), C.uint(unsafe.Sizeof(sockaddr)))
	if err != nil {
		C.close(fd)
		return nil, fmt.Errorf("failed to bind socket address: %w", err.(syscall.Errno))
	}

	return &socket{fd: int(fd)}, nil
}

func addressToSockaddr(addr *Address) C.sockaddr_ax25 {
	return C.sockaddr_ax25{
		sax25_family: C.AF_AX25,
		sax25_call:   addressToAX25(addr),
		sax25_ndigis: 0,
	}
}

func addressToAX25(addr *Address) C.ax25_address {
	var ax25Call C.ax25_address
	for i := 0; i < 6; i++ {
		ax25Call.ax25_call[i] = C.char(addr.Call[i] << 1)
	}
	ax25Call.ax25_call[6] = C.char(((addr.SSID + '0') << 1) & 0x1E)
	return ax25Call
}
