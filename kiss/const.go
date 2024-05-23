package kiss

const (
	// MaxDataSize is a maximum data blob size in one frame.
	MaxDataSize = 2123

	// MaxPDUSize is a maximum protocol data unit (frame) size in bytes.
	MaxPDUSize = MaxDataSize + 3
)

const (
	fend  = 0xc0
	fesc  = 0xdb
	tfend = 0xdc
	tfesc = 0xdd
)
