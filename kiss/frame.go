package kiss

// FrameType represents the type of frame that can be sent or received through KISS.
type FrameType uint8

const (
	// FrameTypeCommand indicates a command frame.
	FrameTypeCommand FrameType = 0x0
)

// Frame represents a single frame of data sent or received through KISS.
type Frame struct {
	Port    uint8
	Command FrameType
	Data    []byte
}
