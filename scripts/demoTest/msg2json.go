package main

import (
	"fmt"
	"io"
	"os"

	"github.com/qrtz/nativemessaging"
)

func main() {
	decoder := nativemessaging.NewNativeJSONDecoder(os.Stdin)

	var msg message

	err := decoder.Decode(&msg)
	if err != nil {
		if err == io.EOF {
			// exit
			return
		}
	}
	fmt.Println("Errors: ", msg.Errors)
	fmt.Println("Resp: ", msg.Resp)
	fmt.Println("Stdout: ", msg.Stdout)
	fmt.Println("Stdout: ", msg.Stdout)
	fmt.Println("\n\n\n")
}

type message struct {
	Errors []string `json:"errors"`
	Resp   []string `json:"resp"`
	Stdout string   `json:"stdout"`
}
