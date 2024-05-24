package kiss

import "io"

// NewEncoder creates a new Encoder that will use the provided writer and port for all subsequent Write calls.
func NewEncoder(w io.Writer, port uint8) *Encoder {
	return &Encoder{
		port:   port,
		writer: NewFrameWriter(w),
	}
}

// Encoder provides an interface for encoding data into KISS frames.
type Encoder struct {
	port   uint8
	writer *FrameWriter
}

// Write encodes data into KISS protocol format and writes it to underlying io.Writer.
func (e *Encoder) Write(data []byte) (n int, err error) {
	for segment := 0; segment < len(data); segment += MaxDataSize {
		segmentEnd := segment + MaxDataSize
		if segmentEnd > len(data) {
			segmentEnd = len(data)
		}

		frame := Frame{
			Port:    e.port,
			Command: FrameTypeCommand,
			Data:    data[segment:segmentEnd],
		}
		err = e.writer.Write(&frame)
		if err != nil {
			return segment, err
		}
	}
	return len(data), nil
}
