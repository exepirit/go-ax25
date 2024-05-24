package kiss

const (
	// MaxDataSize is the maximum size of data blob that can be sent or received in one frame.
	MaxDataSize = 2123

	// MaxPDUSize is a maximum protocol data unit (frame) size in bytes.
	MaxPDUSize = MaxDataSize + 3
)
