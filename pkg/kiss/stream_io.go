package kiss

import "io"

func NewEncoder(w io.Writer, port uint8) *Encoder {
	return &Encoder{
		port:   port,
		writer: NewFrameWriter(w),
	}
}

type Encoder struct {
	port   uint8
	writer *FrameWriter
}

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
