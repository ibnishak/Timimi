package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

var byteOrder binary.ByteOrder = binary.LittleEndian

func receive(reader io.Reader) ([]byte, error) {
	// Read message length in native byte order
	var length uint32
	if err := binary.Read(reader, byteOrder, &length); err != nil {
		return nil, err
	}

	// Return if no message
	if length == 0 {
		return nil, nil
	}

	// Read message body
	received := make([]byte, length)
	if n, err := io.ReadFull(reader, received); err != nil || n != len(received) {
		return nil, err
	}
	return received, nil
}

func main() {
	b, err := receive(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
