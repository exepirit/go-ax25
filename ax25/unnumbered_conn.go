package ax25

import (
	"errors"
	"net"
	"time"
)

func DialUnnumbered(localAddr *Address, remoteAddr *Address) (*UnnumberedConn, error) {
	if localAddr == nil {
		return nil, errors.New("local address must be defined")
	}

	sock, err := openUnnumberedSocket(localAddr)
	return &UnnumberedConn{
		localAddr:  localAddr,
		remoteAddr: remoteAddr,
		sock:       sock,
	}, err
}

type UnnumberedConn struct {
	localAddr  *Address
	remoteAddr *Address
	sock       *socket
}

func (conn *UnnumberedConn) ReadFrom(packet []byte) (n int, addr net.Addr, err error) {
	return conn.sock.recvFrom(packet)
}

func (conn *UnnumberedConn) Read(data []byte) (n int, err error) {
	if conn.remoteAddr == nil {
		return 0, errors.New("destination address is not specified")
	}

	var remote *Address
	for {
		n, remote, err = conn.sock.recvFrom(data)
		if err != nil {
			return 0, err
		}

		if conn.remoteAddr.Equal(remote) {
			return
		}
	}
}

func (conn *UnnumberedConn) WriteTo(packet []byte, addr net.Addr) (n int, err error) {
	return conn.sock.sendTo(packet, addr.(*Address))
}

func (conn *UnnumberedConn) Write(data []byte) (n int, err error) {
	if conn.remoteAddr == nil {
		return 0, errors.New("destination address is not specified")
	}
	return conn.sock.sendTo(data, conn.remoteAddr)
}

func (conn *UnnumberedConn) Close() error {
	return conn.sock.close()
}

func (conn *UnnumberedConn) LocalAddr() net.Addr {
	return conn.localAddr
}

func (conn *UnnumberedConn) SetDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (conn *UnnumberedConn) SetReadDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (conn *UnnumberedConn) SetWriteDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}
