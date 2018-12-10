package main

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var byteOrder binary.ByteOrder = binary.LittleEndian

func Receive(reader io.Reader) ([]byte, error) {
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

type response struct {
	Content string `json:"content"`
}

func main() {

	msg, err := Receive(os.Stdin)
	if err != nil {
		panic(err)
	}
	var res response
	err = json.Unmarshal([]byte(msg), &res)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(res.Content)
	err = ioutil.WriteFile("test.txt", []byte(res.Content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
