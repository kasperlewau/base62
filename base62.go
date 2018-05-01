// Package base62 converts integers into byteslices of base62 data
// and back again.
package base62

import "bytes"
import "strings"
import "unsafe"

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var balphabet = []byte(alphabet)

// Encode encodes the given int and returns a byteslice
func Encode(n int) []byte {
	if n == 0 {
		return balphabet[:1]
	}

	b := make([]byte, 0)
	length := len(alphabet)

	for n > 0 {
		result := n / length
		remainder := n % length
		b = append(b, alphabet[remainder])
		n = result
	}

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// EncodeString encodes the given int and returns a string
func EncodeString(n int) string {
	if n == 0 {
		return string(alphabet[0])
	}

	b := make([]byte, 0)
	length := len(alphabet)

	for n > 0 {
		result := n / length
		remainder := n % length
		b = append(b, alphabet[remainder])
		n = result
	}

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// EncodeStringUnsafe encodes the given int and returns a string using unsafe tricks
func EncodeStringUnsafe(n int) string {
	if n == 0 {
		return *(*string)(unsafe.Pointer(&balphabet[0]))
	}

	b := make([]byte, 0)
	length := len(alphabet)

	for n > 0 {
		result := n / length
		remainder := n % length
		b = append(b, balphabet[remainder])
		n = result
	}

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return *(*string)(unsafe.Pointer(&b))
}

// Decode returns a numeric representation of the given base62 byteslice
func Decode(b []byte) int {
	n, length := 0, len(alphabet)

	for _, c := range b {
		i := bytes.IndexByte(balphabet, c)
		if i < 0 {
			return n
		}
		n = length*n + i
	}

	return n
}

// DecodeString returns a numeric representation of the given base62 string
func DecodeString(s string) int {
	n, length := 0, len(alphabet)

	for _, c := range s {
		i := strings.IndexRune(alphabet, c)
		if i < 0 {
			return n
		}
		n = length*n + i
	}

	return n
}
