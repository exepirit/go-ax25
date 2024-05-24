package ax25

// UnnumberedPacket represents an AX.25 UI packet.
type UnnumberedPacket struct {
	Address PacketAddress
	PID     ProtocolID
	Info    []byte
}

// PacketAddress represents the packet two-way address.
type PacketAddress struct {
	Destination Address
	Source      Address
}

// ProtocolID represents a type of protocol in the PID field.
type ProtocolID byte

const (
	// ProtocolNoLayer3 specifies that no Layer 3 protocol is used in this frame.
	ProtocolNoLayer3 ProtocolID = 0xF0
)
