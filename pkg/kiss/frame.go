package kiss

const (
	DataFrameCommand = 0x0
)

// Frame implements simple KISS TNC frame.
type Frame struct {
	Port    uint8
	Command uint8
	Data    []byte
}
