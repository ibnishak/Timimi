package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var byteOrder binary.ByteOrder = binary.LittleEndian
var wg sync.WaitGroup

// done := make(chan bool)
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

func main() {

	msg, err := receive(os.Stdin)
	if err != nil {
		// panic(err)
		send("StdIn failed with error", err.Error())
	}
	var data indata
	err = json.Unmarshal([]byte(msg), &data)
	if err != nil {
		send("UnMarshall Failed with error ", err.Error())
		// panic(err)
	}
	// send(data.Bstrategy, data.Tnow)
	if data.Path != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = ioutil.WriteFile(data.Path, []byte(data.Content), 0666)
			if err != nil {
				// log.Fatal(err)
				send("Save failed with error", err.Error())
			}
			send("Saved Successfully to ", data.Path)
		}()
	}
	if data.Backup == "yes" {
		wg.Add(1)
		go backup(data)
	}
	if data.Exec == "yes" {
		// send("Point", "C")
		efinal := filepath.Join(os.Getenv("HOME"), ".timimi", data.Escript)
		cmd := exec.Command(efinal, data.Eparam)
		if data.Estdin != "" {
			cmd.Stdin = strings.NewReader(data.Estdin)
		}
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			send("Action Script Error: ", err.Error())
		}
		send("STDOUT: ", out.String())
		// fmt.Printf("in all caps: %q\n", out.String())
	}
	wg.Wait()
}

func backup(data indata) {
	defer wg.Done()
	mid := data.MessageID

	var tdir string
	if data.Bpath != "" {
		if path.IsAbs(data.Bpath) {
			tdir = data.Bpath
		} else {
			tdir = path.Join(path.Dir(data.Path), data.Bpath)
		}
	} else {
		tdir = path.Dir(data.Path)
	}
	var tfile = filenameWithoutExtension(path.Base(data.Path))
	var ext = path.Ext(data.Path)
	save := false
	var tfinal string
	// TOWER OF HANOI LOGIC
	if data.Bstrategy == "toh" {
		// var l, r int = data.Tohlevel, data.Tohrecent
		l, err := strconv.Atoi(data.Tohlevel)
		if err != nil {
			send("Error during conversion of tohlevel to int: ", err.Error())
			return
		}
		r, err := strconv.Atoi(data.Tohrecent)
		if err != nil {
			send("Error during conversion of tohrecent to int: ", err.Error())
			return
		}
		// TOH SNAPSHOT.LOOP
		for n := float64(l); n >= 3; n-- {
			if mid%8 != 0 {
				break
			}
			p := int(math.Pow(2, n))
			if mid%p == 0 {
				tfinal = fmt.Sprintf("%s-%d%s", tfile, p, ext)
				err := ioutil.WriteFile(path.Join(tdir, tfinal), []byte(data.Content), 0666)
				if err != nil {
					// log.Fatal(err)
					send("TOH Snapshot backup failed with error", err.Error())
				}
				// fmt.Printf("Saved to %s\n", tfinal)
				save = true
				break
			}
		} // TOH SNAPSHOT LOOP ENDS

		// TOH RECENT BACKUPS
		if save != true {
			if mid < r {
				tfinal = fmt.Sprintf("%s-A%d%s", tfile, mid, ext)
				err := ioutil.WriteFile(path.Join(tdir, tfinal), []byte(data.Content), 0666)
				if err != nil {
					// log.Fatal(err)
					send("Error: TOH recent backups failed", err.Error())
				}

				// fmt.Printf("Saved to %s\n", tfinal)
			} else {
				c := mid % r
				tfinal = fmt.Sprintf("%s-A%d%s", tfile, c, ext)
				err := ioutil.WriteFile(path.Join(tdir, tfinal), []byte(data.Content), 0666)
				if err != nil {
					send("Error: TOH recent backups failed with error: ", err.Error())
				}
				// fmt.Printf("Saved to %s\n", tfinal)
			}
		} // TOH RECENT BACKUPS END
		//TOH LOGIC ENDS
	} else if data.Bstrategy == "psave" {
		pint, err := strconv.Atoi(data.Psint)
		if err != nil {
			send("Error during conversion of Psint to int: ", err.Error())
			return
		}
		if mid%pint == 0 {

			tfinal = fmt.Sprintf("%s-%s%s", tfile, buildFileName(), ext)
			err := ioutil.WriteFile(path.Join(tdir, tfinal), []byte(data.Content), 0666)
			if err != nil {
				// log.Fatal(err)
				send("Error: Per save backups failed", err.Error())
			}
			//fmt.Printf("Saved to %s\n", tfinal)
		}
		//PERSAVE BACKUP LOGIC ENDS
	} else if data.Bstrategy == "timed" {
		if data.TBackup == "true" {
			tfinal = fmt.Sprintf("%s-%s%s", tfile, buildFileName(), ext)
			err := ioutil.WriteFile(path.Join(tdir, tfinal), []byte(data.Content), 0666)
			if err != nil {
				// log.Fatal(err)
				send("Error: Timed backups failed", err.Error())
			}
		} else {
			return
		}

	}
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func buildFileName() string {
	return time.Now().Format("2006-01-02-15-04-05")
}

func send(res, info string) {
	reply := fmt.Sprintf(`{ "content":"%s %s"}`, res, info)
	post([]byte(reply), os.Stdout)
}
