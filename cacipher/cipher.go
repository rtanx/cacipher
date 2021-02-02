package cacipher

import (
	"io"
	"strings"
)

//  cipher holds the requirements for caesar-cipher encode/decode
type cipher struct {
	text  string    // ecoded/pre-encoded text string
	shift int       // number of shifts
	r     io.Reader // std reader interface
}

// Encode encode a plaintext to the caesar-chiper encryption
func Encode(plaintxt string, shift uint) string {
	c := cipher{
		text:  plaintxt,
		shift: int(shift),
		r:     strings.NewReader(plaintxt),
	}
	ciphertxt := new(strings.Builder)
	io.Copy(ciphertxt, c)
	return ciphertxt.String()
}

// Decode decode a caesar-cipher encryption to the original plaintext
func Decode(ciphertxt string, shift uint) string {
	c := cipher{
		text:  ciphertxt,
		shift: -int(shift),
		r:     strings.NewReader(ciphertxt),
	}
	plaintxt := new(strings.Builder)
	io.Copy(plaintxt, c)
	return plaintxt.String()
}

func (c cipher) Read(p []byte) (n int, err error) {
	n, err = c.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = c.shifter(p[i])
	}
	return
}
func (c cipher) shifter(b byte) byte {
	shift := (c.shift%26 + 26) % 26
	var char byte
	switch {
	case 'a' <= b && b <= 'z':
		char = 'a'
	case 'A' <= b && b <= 'Z':
		char = 'A'
	default:
		return b
	}
	return (char + ((b-char)+byte(shift))%26)
}
