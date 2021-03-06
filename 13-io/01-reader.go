package main

import (
	"fmt"
	"io"
	"strings"
)

func firstMain() { //rename to main for its execution

	reader := strings.NewReader("Hello, Reader!")

	bytes := make([]byte, 8)

	for {
		n, err := reader.Read(bytes) //consumes 8 bytes at a time
		fmt.Printf("n = %v err = %v b = %v\n", n, err, bytes)
		fmt.Printf("b[:n] = %q\n", bytes[:n])

		if err == io.EOF {
			break
		}
	}
}
