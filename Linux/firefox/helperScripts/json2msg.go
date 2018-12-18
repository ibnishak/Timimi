package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/qrtz/nativemessaging"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var msg message
	err = json.Unmarshal(data, &msg)
	if err != nil {
		log.Fatal(err)
	}
	encoder := nativemessaging.NewNativeJSONEncoder(os.Stdout)
	err = encoder.Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
}

type message struct {
	Content   string `json:"content"`
	Path      string `json:"path"`
	Backup    string `json:"backup"`
	Bpath     string `json:"bpath"`
	Bstrategy string `json:"bstrategy"`
	MessageID int    `json:"messageId"`
	Tohrecent string `json:"tohrecent"`
	Tohlevel  string `json:"tohlevel"`
	Psint     string `json:"psint"`
	Exec      string `json:"exec"`
	Escript   string `json:"escript"`
	Eparam    string `json:"eparam"`
	Estdin    string `json:"estdin"`
	TBackup   string `json:"tbackup"`
}
