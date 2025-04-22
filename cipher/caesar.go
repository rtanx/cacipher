package cipher

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// Caesar struct holds all the necessary components for encryption/decryption
type Caesar struct {
	text    string    // text string to be encrypted/decrypted
	shift   int       // number of shifts
	r       io.Reader // input reader interface
	w       io.Writer // output writer interface
	decrypt bool      // flag to determine operation (encrypt/decrypt)
}

// NewCaesar creates and returns a new Caesar instance
func NewCaesar(shift int, reader io.Reader, writer io.Writer, decrypt bool) *Caesar {
	return &Caesar{
		text:    "",
		shift:   shift,
		r:       reader,
		w:       writer,
		decrypt: decrypt,
	}
}

// transform applies the Caesar cipher transformation on the text
func (c *Caesar) transform() string {
	var result strings.Builder
	shift := c.shift

	// If decrypting, negate the shift value
	if c.decrypt {
		shift = -shift
	}

	// Make sure shift is within 0-25 range (after applying modulo)
	shift = ((shift % 26) + 26) % 26

	for _, char := range c.text {
		if unicode.IsLetter(char) {
			// Determine if uppercase or lowercase
			base := 'a'
			if unicode.IsUpper(char) {
				base = 'A'
			}

			// Apply the shift and wrap around the alphabet
			shifted := (int(char) - int(base) + shift) % 26
			result.WriteRune(rune(shifted + int(base)))
		} else {
			// Non-alphabetic characters remain unchanged
			result.WriteRune(char)
		}
	}

	return result.String()
}

// Transform reads from input, performs cipher operation, and writes to output
func (c *Caesar) Transform() error {
	scanner := bufio.NewScanner(c.r)
	for scanner.Scan() {
		c.text = scanner.Text()
		result := c.transform()
		if _, err := fmt.Fprintln(c.w, result); err != nil {
			return err
		}
	}
	return scanner.Err()
}

// TransformText transforms the provided text and writes to output
func (c *Caesar) TransformText(text string) error {
	c.text = text
	result := c.transform()
	_, err := fmt.Fprintln(c.w, result)
	return err
}

func (c *Caesar) SetReader(reader io.Reader) {
	c.r = reader
}
