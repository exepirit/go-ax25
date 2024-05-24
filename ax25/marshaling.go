package ax25

import (
	"io"
)

// NewPacketWriter creates a new instance of PacketEncoder for writing AX.25 packets to an io.Writer interface.
// The `mtu` (Maximum Transfer Unit) parameter determines the maximum size of the packet information field.
func NewPacketWriter(w io.Writer, mtu uint16) *PacketEncoder {
	return &PacketEncoder{
		buffer: newBytesBuffer(int(14 + 2 + 1 + mtu)), // max possible frame size
		writer: w,
	}
}

// PacketEncoder encodes and writes AX.25 packets to an output stream.
type PacketEncoder struct {
	buffer bytesBuffer
	writer io.Writer
}

// Write writes a packet to the output stream.
func (penc *PacketEncoder) Write(p *Packet) error {
	penc.buffer.reset()
	for _, b := range p.Address.Destination.Call {
		penc.buffer.writeByte(b << 1)
	}
	penc.buffer.writeByte(p.Address.Destination.SSID<<1 | 0xE0)
	for _, b := range p.Address.Source.Call {
		penc.buffer.writeByte(b << 1)
	}
	penc.buffer.writeByte(p.Address.Source.SSID<<1 | 0x61)
	penc.writeControl(&p.Control)
	penc.buffer.writeByte(byte(p.PID))
	penc.buffer.write(p.Info)

	_, err := penc.writer.Write(penc.buffer.bytes())
	return err
}

func (penc *PacketEncoder) writeControl(cf *ControlData) {
	var b byte
	switch cf.Type {
	case PacketTypeUnnumbered:
		b = 0x3
		if cf.IsFinal {
			b |= 0x10
		}
	default:
		panic("unknown packet type")
	}
	penc.buffer.writeByte(b)
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
