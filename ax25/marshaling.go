package ax25

import (
	"io"
)

// NewFrameEncoder creates a new instance of FrameEncoder for writing AX.25 frames to an io.Writer interface.
// The `mtu` (Maximum Transfer Unit) parameter determines the maximum size of the frame information field.
func NewFrameEncoder(w io.Writer, mtu uint16) *FrameEncoder {
	return &FrameEncoder{
		buffer: newBytesBuffer(int(14 + 2 + 1 + mtu)), // max possible frame size
		writer: w,
	}
}

// FrameEncoder encodes and writes AX.25 frames to an output stream.
type FrameEncoder struct {
	buffer bytesBuffer
	writer io.Writer
}

// Write writes a frame to the output stream.
func (penc *FrameEncoder) Write(frame *Frame) error {
	penc.buffer.reset()
	for _, b := range frame.Address.Destination.Call {
		penc.buffer.writeByte(b << 1)
	}
	penc.buffer.writeByte(frame.Address.Destination.SSID<<1 | 0xE0)
	for _, b := range frame.Address.Source.Call {
		penc.buffer.writeByte(b << 1)
	}
	penc.buffer.writeByte(frame.Address.Source.SSID<<1 | 0x61)
	penc.writeControl(&frame.Control)
	penc.buffer.writeByte(byte(frame.PID))
	penc.buffer.write(frame.Info)

	_, err := penc.writer.Write(penc.buffer.bytes())
	return err
}

func (penc *FrameEncoder) writeControl(cf *ControlData) {
	var b byte
	switch cf.Type {
	case FrameTypeUnnumbered:
		b = 0x3
		if cf.IsFinal {
			b |= 0x10
		}
	default:
		panic("unknown frame type")
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
