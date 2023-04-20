package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)

}

func Reverse(s string) (string, error) {

	fmt.Printf("input: %q\n", s)

	// In second version, we found the bug happened when the input string will not be a utf8 string
	// So we add a validating method for input string.
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}

	// In first version used byte method to convert characters of string from input,
	// and it had a bug with the storage units are not the same.
	// Some characters which are defined in the ascii code table used 1 byte, otherwise may have used 3 or 4 bytes.
	// So we change the converting method from byte to rune to guarantee each character have the same byte length of string.
	// b := []byte(s)
	b := []rune(s)

	fmt.Printf("runes: %q\n", b)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}
