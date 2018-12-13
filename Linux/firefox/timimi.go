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

type indata struct {
	Content string `json:"content"`
	Path    string `json:"path"`
}

func main() {

	msg, err := receive(os.Stdin)
	if err != nil {
		panic(err)
	}
	var data indata
	err = json.Unmarshal([]byte(msg), &data)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(data.Content)
	err = ioutil.WriteFile(data.Path, []byte(data.Content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
