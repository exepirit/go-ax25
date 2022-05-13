package kiss

import (
	"errors"
	"io"
)

const MaxFrameSize = 2500

// NewFrameWriter wraps io.Writer to FrameWriter.
func NewFrameWriter(w io.Writer) *FrameWriter {
	return &FrameWriter{
		w: w,
	}
}

// FrameWriter writes frames to io.Writer.
type FrameWriter struct {
	w io.Writer
}

func (writer FrameWriter) Write(frame Frame) error {
	escaped := Escape(frame.Data)
	buf := make([]byte, len(escaped)+3)
	buf[0], buf[len(buf)-1] = fend, fend
	buf[1] = frame.Command<<4 | frame.Port
	copy(buf[2:len(buf)-1], escaped)

	_, err := writer.w.Write(buf)
	return err
}

// NewFrameReader wraps io.Reader to FrameReader.
func NewFrameReader(r io.Reader) *FrameReader {
	return &FrameReader{
		r: r,
	}
}

// FrameReader reads frames from io.Reader.
type FrameReader struct {
	r io.Reader
}

func (reader FrameReader) Read() (Frame, error) {
	frameData := make([]byte, MaxFrameSize)
	n, err := reader.r.Read(frameData)
	if err != nil {
		return Frame{}, err
	}
	frameData = frameData[:n]

	if len(frameData) < 3 {
		return Frame{}, io.ErrUnexpectedEOF
	}

	if frameData[0] != fend || frameData[len(frameData)-1] != fend {
		return Frame{}, errors.New("data is not valid TNC frame")
	}

	frame := Frame{
		Port:    frameData[1] & 0xf,
		Command: frameData[1] >> 4 & 0xf,
	}
	frame.Data, err = Unescape(frameData[2 : len(frameData)-1])
	return frame, err
}
