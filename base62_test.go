package base62_test

import (
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
