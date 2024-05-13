package kiss

type FrameType uint8

const (
	FrameTypeCommand FrameType = 0x0
)

// Frame implements simple KISS TNC frame.
type Frame struct {
	Port    uint8
	Command FrameType
	Data    []byte
}
