package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var byteOrder binary.ByteOrder = binary.LittleEndian

func post(msg []byte, writer io.Writer) error {
	// Post message length in native byte order
	header := make([]byte, 4)
	byteOrder.PutUint32(header, (uint32)(len(msg)))
	if n, err := writer.Write(header); err != nil || n != len(header) {
		return err
	}

	// Post message body
	if n, err := writer.Write(msg); err != nil || n != len(msg) {
		return err
	}
	return nil
}

func main() {
	file := os.Args[1]
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	post(b, os.Stdout)
}
