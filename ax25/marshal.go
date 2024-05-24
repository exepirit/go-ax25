package ax25

import (
	"io"
)

// NewPacketWriter creates a new instance of PacketWriter for writing AX.25 packets to an io.Writer interface.
// The `mtu` (Maximum Transfer Unit) parameter determines the maximum size of the packet information field.
func NewPacketWriter(w io.Writer, mtu uint16) *PacketWriter {
	return &PacketWriter{
		buffer: newBytesBuffer(int(14 + 2 + 1 + mtu)), // max possible frame size
		writer: w,
	}
}

// PacketWriter writes AX.25 packets to an output stream.
type PacketWriter struct {
	buffer bytesBuffer
	writer io.Writer
}

// WriteUnnumbered writes an UI packet to the output stream.
func (pw *PacketWriter) WriteUnnumbered(p *UnnumberedPacket) error {
	pw.buffer.reset()
	for _, b := range p.Address.Destination.Call {
		pw.buffer.writeByte(b << 1)
	}
	pw.buffer.writeByte(p.Address.Destination.SSID<<1 | 0xE0)
	for _, b := range p.Address.Source.Call {
		pw.buffer.writeByte(b << 1)
	}
	pw.buffer.writeByte(p.Address.Source.SSID<<1 | 0x61)
	pw.buffer.writeByte(0x3) // Control field: 00000011
	pw.buffer.writeByte(byte(p.PID))
	pw.buffer.write(p.Info)

	_, err := pw.writer.Write(pw.buffer.bytes())
	return err
}

func newBytesBuffer(size int) bytesBuffer {
	return bytesBuffer{
		buf: make([]byte, size), // max possible frame size
	}
}

type bytesBuffer struct {
	buf []byte
	ptr uint16
}

func (fb *bytesBuffer) writeByte(b byte) {
	fb.buf[fb.ptr] = b
	fb.ptr++
}

func (fb *bytesBuffer) write(data []byte) {
	for _, b := range data {
		fb.writeByte(b)
	}
}

func (fb *bytesBuffer) reset() {
	fb.ptr = 0
}

func (fb *bytesBuffer) bytes() []byte {
	return fb.buf[:fb.ptr]
}
