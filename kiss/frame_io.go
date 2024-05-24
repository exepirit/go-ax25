package kiss

import (
	"bufio"
	"io"
)

// NewFrameWriter creates a new writer that will write KISS frames to the provided io.Writer.
func NewFrameWriter(w io.Writer) *FrameWriter {
	return &FrameWriter{
		buf: bufio.NewWriterSize(w, MaxPDUSize),
	}
}

const (
	fend  = 0xc0
	fesc  = 0xdb
	tfend = 0xdc
	tfesc = 0xdd
)

// FrameWriter writes frames to io.Writer.
type FrameWriter struct {
	buf *bufio.Writer
}

// Write encodes frame into KISS protocol format and writes to underlying io.Writer through buffer.
func (writer *FrameWriter) Write(frame *Frame) error {
	err := writer.buf.WriteByte(fend)
	if err != nil {
		return err
	}

	err = writer.buf.WriteByte(uint8(frame.Command)<<4 | (frame.Port & 0xf))
	if err != nil {
		return err
	}

	for _, b := range frame.Data {
		switch b {
		case fend:
			_, err = writer.buf.Write([]byte{fesc, tfend})
		case fesc:
			_, err = writer.buf.Write([]byte{fesc, tfesc})
		default:
			err = writer.buf.WriteByte(b)
		}

		if err != nil {
			return err
		}
	}

	err = writer.buf.WriteByte(fend)
	if err != nil {
		return err
	}

	return writer.buf.Flush()
}

// NewFrameReader creates a new FrameReader that will use the provided io.Reader as its source of data.
func NewFrameReader(r io.Reader) *FrameReader {
	return &FrameReader{
		scanner: bufio.NewReaderSize(r, MaxPDUSize),
	}
}

type frameReaderState uint8

const (
	readerWaitFrame frameReaderState = iota
	readerWaitCommand
	readerCommandRead
	readerWaitData
	readerEscapeSeqRead
)

// FrameReader provides an interface for reading KISS frames.
type FrameReader struct {
	scanner *bufio.Reader
	state   frameReaderState
	frame   Frame
}

// Read decodes next frame from I/O into Frame format.
func (reader *FrameReader) Read() (*Frame, error) {
	for {
		b, err := reader.scanner.ReadByte()
		if err != nil {
			return nil, err
		}

		if reader.handleByte(b) {
			return &reader.frame, nil
		}
	}
}

func (reader *FrameReader) handleByte(b byte) (done bool) {
	switch reader.state {
	case readerWaitFrame:
		if b == fend {
			reader.state = readerWaitCommand
		}
	case readerWaitCommand:
		reader.frame.Port = b & 0xf
		reader.frame.Command = FrameType(b >> 4 & 0xf)
		reader.frame.Data = make([]byte, 0, MaxDataSize)
		reader.state = readerCommandRead

	case readerCommandRead:
		if b == fesc {
			reader.state = readerEscapeSeqRead
			return
		}
		reader.frame.Data = append(reader.frame.Data, b)
		reader.state = readerWaitData

	case readerEscapeSeqRead:
		switch b {
		case tfend:
			b = fend
		case tfesc:
			b = fesc
		}
		reader.frame.Data = append(reader.frame.Data, b)
		reader.state = readerWaitData

	case readerWaitData:
		if b == fesc {
			reader.state = readerEscapeSeqRead
			return
		}
		if b == fend {
			reader.state = readerWaitFrame
			done = true
			return
		}
		reader.frame.Data = append(reader.frame.Data, b)

	default:
		panic("unknown state")
	}
	return
}
