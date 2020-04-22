package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	TBackup   string `json:"tbackup"`
}

func main() {
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "Timimi", log.LstdFlags)

	msg, err := receive(os.Stdin)
	if err != nil {
		notifier(err.Error(), logger)
		os.Exit(1)
	}

	data, err := unmarshdata(msg)
	if err != nil {
		notifier(err.Error(), logger)
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
	default:
		bname, bstatus = "", false
	}

	if bstatus {
		bpath := setbackuppath(data.Bpath, data.Path)
		err = ensuredir(bpath)
		if err != nil {
			notifier(err.Error(), logger)
		}
		wg.Add(1)
		go savetw(filepath.Join(bpath, bname), data.Content)
	}
	wg.Wait()
	reply, _ := json.Marshal("Timimi Host exits")
	post(reply, os.Stdout)
}

func savetw(path, content string) error {
	defer wg.Done()
	err := ioutil.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err
	}
	return nil
}
