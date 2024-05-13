package kiss

import (
	"bytes"
	"testing"
)

func TestReader_Unescaped(t *testing.T) {
	data := []byte{
		0xc0, 0x24, 0x23, 0xc0,
		0xc0, 0x00, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x2d, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0xc0,
	}
	buf := bytes.NewBuffer(data)
	reader := NewFrameReader(buf)

	frame, err := reader.Read()
	if err != nil {
		t.Fatal(err)
	}

	frame, err = reader.Read()
	if err != nil {
		t.Fatal(err)
	}

	if string(frame.Data) != "HELLO-WORLD" {
		t.Logf("expected = \"HELLO-WORLD\", got = %q", frame.Data)
		t.Fail()
	}
}
