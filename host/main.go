package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

type indata struct {
	Content   string `json:"content"`
	Path      string `json:"path"`
	Backup    string `json:"backup"`
	Bpath     string `json:"bpath"`
	Bstrategy string `json:"bstrategy"`
	MessageID int    `json:"messageId"`
	Tohrecent string `json:"tohrecent"`
	Tohlevel  string `json:"tohlevel"`
	Psint     string `json:"psint"`
	Fint      string `json:"fint"`
	TBackup   string `json:"tbackup"`
}

func main() {

	msg, err := receive(os.Stdin)
	if err != nil {
		notifier()
		logerr("Error while receiving stdin", err.Error())
		reply, _ := json.Marshal("Timimi Host Error: Error while receiving stdin")
		post(reply, os.Stdout)
		os.Exit(1)
	}

	data, err := unmarshdata(msg)
	if err != nil {
		notifier()
		logerr("Error while unmarshalling stdin", err.Error())
		reply, _ := json.Marshal("Timimi Host Error: Error while unmarshalling stdin")
		post(reply, os.Stdout)
		os.Exit(1)
	}

	wg.Add(1)
	go savetw(data.Path, data.Content)

	var bname string
	var bstatus bool

	switch data.Bstrategy {
	case "toh":
		l, _ := strconv.Atoi(data.Tohlevel)
		r, _ := strconv.Atoi(data.Tohrecent)
		bname, bstatus = toh(data.Path, l, r, data.MessageID)
	case "timed":
		bname, bstatus = timed(data.TBackup, data.Path)
	case "psave":
		bname, bstatus = psave(data.Psint, data.Path, data.MessageID)
	case "fifo":
		bname, bstatus = fifo(data.Fint, data.MessageID, data.Path)
	default:
		bname, bstatus = "", false
	}

	if bstatus {
		bpath := setbackuppath(data.Bpath, data.Path)
		err = ensuredir(bpath)
		if err != nil {
			logerr("Error while creating backup dir", err.Error())
		}
		wg.Add(1)
		go savetw(filepath.Join(bpath, bname), data.Content)
	}
	wg.Wait()
	reply, _ := json.Marshal(fmt.Sprintf("Timimi host successfully saved file to %s. Backup status:%t. Backup path: '%s'", data.Path, bstatus, bname))
	post(reply, os.Stdout)
}

func savetw(path, content string) {
	defer wg.Done()
	err := ioutil.WriteFile(path, []byte(content), 0666)
	if err != nil {
		notifier()
		logerr("Error while writing file", err.Error())
		reply, _ := json.Marshal("Timimi Host Error: Error while writing file")
		post(reply, os.Stdout)
	}
}
