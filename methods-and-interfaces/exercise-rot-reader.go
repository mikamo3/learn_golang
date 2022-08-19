package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyz"
	encryptionKey = 13
	table         = alphabet + alphabet
)

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < len(b); i++ {
		b[i] = r.rotate(b[i], encryptionKey)
	}
	return n, err
}
func (r *rot13Reader) rotate(b byte, key int) byte {
	ch := rune(b)
	if !unicode.IsLetter(ch) {
		return b
	}

	key = key % len(alphabet)

	isUpper := false
	if unicode.IsUpper(ch) {
		isUpper = true
		ch = unicode.ToLower(ch)
	}

	idx := byte(ch) - 'a' + byte(key)
	rotatedCh := rune(table[idx])
	if isUpper {
		rotatedCh = unicode.ToUpper(rotatedCh)
	}
	return byte(rotatedCh)
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
