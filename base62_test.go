package base62_test

import (
	"testing"

	"github.com/kasperlewau/shorten/base62"
)

func TestEncodeDecode(t *testing.T) {
	for i := 0; i < 90; i++ {
		ui := uint64(i)
		enc := base62.Encode(ui)
		dec := base62.Decode(enc)
		if dec != ui {
			t.Errorf("Expected decoding of encoded value %v to equal source %v. Got %v\n", enc, ui, dec)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		base62.Encode(uint64(i))
	}
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()
	buf := []byte("XYZ123")
	for i := 0; i < b.N; i++ {
		base62.Decode(buf)
	}
}
