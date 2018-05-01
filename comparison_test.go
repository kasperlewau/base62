package base62_test

import (
	"testing"

	kl "github.com/kasperlewau/base62"
	"github.com/marksalpeter/token"
	pilu "github.com/pilu/go-base62"
)

func BenchmarkEncoders(b *testing.B) {
	b.Run("kasperlewau/bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			kl.Encode(i)
		}
	})
	b.Run("kasperlewau/string", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			kl.EncodeString(i)
		}
	})
	b.Run("pilu/go-base62", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			pilu.Encode(i)
		}
	})
	b.Run("marksalpeter/token", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			token.New().Encode()
		}
	})
}

func BenchmarkDecoders(b *testing.B) {
	b.Run("kasperlewau/bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			kl.Decode([]byte("foobar"))
		}
	})
	b.Run("kasperlewau/string", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			kl.DecodeString("foobar")
		}
	})
	b.Run("pilu/go-base62", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			pilu.Decode("foobar")
		}
	})
	b.Run("marksalpeter/token", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			token.Decode("foobar")
		}
	})
}
