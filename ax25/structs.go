package ax25

type PacketAddress struct {
	Destination Address
	Source      Address
}

type UnnumberedPacket struct {
	Address PacketAddress
	PID     ProtocolID
	Info    []byte
}

type ProtocolID byte

const (
	ProtocolNoLayer3 ProtocolID = 0xF0
)
