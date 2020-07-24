package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// Comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	// Set the buffer with it 0 value
	var buff bytes.Buffer
	// Use this remainder for get the number of digits behind
	remainder := len(s) % 3

	for i, v := range s {
		// Print just the value if isn't a multiple of 3 after the reminder
		if i < remainder || !((i-remainder)%3 == 0 && i != 0) {
			_, _ = fmt.Fprintf(&buff, "%c", v)
			continue
		}

		// Write a comma is it, then just print the rune
		buff.WriteByte(',')
		_, _ = fmt.Fprintf(&buff, "%c", v)
	}

	return buff.String()
}
