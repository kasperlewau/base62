package base62_test

import (
	"bytes"
	"testing"

	"github.com/kasperlewau/base62"
)

func TestEncodeDecode(t *testing.T) {
	for i := 0; i < 90; i++ {
		enc := base62.Encode(i)
		dec := base62.Decode(enc)
		if dec != i {
			t.Errorf("Expected decoding of encoded value %v to equal source %v. Got %v\n", enc, i, dec)
		}
	}
}

func TestEncodeDecodeString(t *testing.T) {
	for i := 0; i < 90; i++ {
		enc := base62.EncodeString(i)
		dec := base62.DecodeString(enc)
		if dec != i {
			t.Errorf("Expected decoding of encoded value %v to equal source %v. Got %v\n", enc, i, dec)
		}
	}
}

// pilu tests
func TestEncode(t *testing.T) {
	if !bytes.Equal([]byte("0"), base62.Encode(0)) {
		t.Error("Expected Encode(0) to equal 0")
	}
	if !bytes.Equal([]byte("1B"), base62.Encode(99)) {
		t.Error("Expected Encode(99) to equal 1B")
	}
}

func TestDecode(t *testing.T) {
	if 0 != base62.DecodeString("0") {
		t.Error("Expected DecodeString('0') to equal 0")
	}
	if 99 != base62.DecodeString("1B") {
		t.Error("Expected DecodeString('1B') to equal 99")
	}
}
