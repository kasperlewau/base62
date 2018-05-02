package base62_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kasperlewau/base62"
)

func TestEncodeDecode(t *testing.T) {
	for i := 0; i < 10000; i++ {
		enc := base62.Encode(i)
		dec := base62.Decode(enc)
		if dec != i {
			t.Errorf("Expected decoding of encoded value %v to equal source %v. Got %v\n", enc, i, dec)
		}
	}
}

func TestEncodeDecodeString(t *testing.T) {
	for i := 0; i < 10000; i++ {
		enc := base62.EncodeString(i)
		dec := base62.DecodeString(enc)
		if dec != i {
			t.Errorf("Expected decoding of encoded value %v to equal source %v. Got %v\n", enc, i, dec)
		}
	}
}

func TestDecodeNegative(t *testing.T) {
	if 0 != base62.Decode([]byte("-1")) {
		t.Error("Expected decoding of -1 to equal 0")
	}
}

func TestDecodeStringNegative(t *testing.T) {
	if 0 != base62.DecodeString("-1") {
		t.Error("Expected decoding of -1 to equal 0")
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

func ExampleDecode() {
	fmt.Printf("%v", base62.Decode([]byte("1B")))
	// Output:
	// 99
}

func ExampleDecodeString() {
	fmt.Println(base62.DecodeString("1B"))
	// Output:
	// 99
}

func ExampleEncode() {
	fmt.Printf("%s\n", base62.Encode(99))
	// Output:
	// 1B
}

func ExampleEncodeString() {
	fmt.Println(base62.EncodeString(99))
	// Output:
	// 1B
}
