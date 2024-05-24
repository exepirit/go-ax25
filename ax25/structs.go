package ax25

// Frame represents an AX.25 frame.
type Frame struct {
	Address FrameAddress
	Control ControlData
	PID     ProtocolID
	Info    []byte
}

// FrameAddress represents the frame two-way address.
type FrameAddress struct {
	Destination Address
	Source      Address
}

// ProtocolID represents a type of underlying protocol.
type ProtocolID byte

const (
	// ProtocolNoLayer3 specifies that no Layer 3 protocol is used in this frame.
	ProtocolNoLayer3 ProtocolID = 0xF0
)

// ControlData represents flow control data.
type ControlData struct {
	Type    FrameType
	IsFinal bool
}

// FrameType represents a type of AX.25 frame.
type FrameType int

const (
	FrameTypeUnnumbered FrameType = iota
)
