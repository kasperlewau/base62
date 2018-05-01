// Package base62 converts integers into byteslices of base62 data
// and back again.
package base62

import "bytes"

var alphabet = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var length = len(alphabet)

// Encode returns the base62 representation of the given number
func Encode(n int) []byte {
	if n == 0 {
		return alphabet[:1]
	}

	b := []byte{}

	for ; n > 0; n /= length {
		b = append([]byte{alphabet[n%length]}, b...)
	}

	return b
}

// Decode returns a numeric representation of the given base62 bytes
func Decode(b []byte) int {
	var n int
	for _, c := range b {
		i := bytes.IndexByte(alphabet, c)
		if i < 0 {
			return n
		}
		n = length*n + i
	}
	return n
}
