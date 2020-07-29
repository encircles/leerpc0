package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterCodec(t *testing.T) {
	RegisterCodec("testCodec", nil)

	codec := GetCodec("testCodec")
	assert.Equal(t, codec, nil)
}

func TestDefaultCodec_Decode(t *testing.T) {
	c := DefaultCodec
	bytes, err := c.Encode([]byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes)
}

func TestDefaultCodec_Encode(t *testing.T) {
	var b []byte
	b = append(b, 0)
	_ = b[0]
}
