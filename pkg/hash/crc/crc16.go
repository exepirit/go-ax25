package crc

func NewCRC16CCIT() CRC16 {
	return NewCRC16(0xFFFF, 0x8408)
}

func NewCRC16(init, polynom uint16) CRC16 {
	return CRC16{
		cur:  init,
		poly: polynom,
	}
}

type CRC16 struct {
	cur  uint16
	poly uint16
}

func (c *CRC16) Value() uint16 {
	return c.cur
}

func (c *CRC16) Update(data []byte) {
	for _, b := range data {
		c.UpdateByte(b)
	}
}

func (c *CRC16) UpdateByte(b byte) {
	for i := 0; i < 8; i++ {
		if (c.cur&0x0001)^(uint16(b)&0x0001) == 1 {
			c.cur = (c.cur >> 1) ^ c.poly
		} else {
			c.cur >>= 1
		}
		b >>= 1
	}
}
