package ax25

// Packet represents an AX.25 packet.
type Packet struct {
	Address PacketAddress
	Control ControlData
	PID     ProtocolID
	Info    []byte
}

// PacketAddress represents the packet two-way address.
type PacketAddress struct {
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
	Type    PacketType
	IsFinal bool
}

// PacketType represents a type of AX.25 packet.
type PacketType int

const (
	PacketTypeUnnumbered PacketType = iota
)
