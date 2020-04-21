package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var byteOrder binary.ByteOrder = binary.LittleEndian
var wg sync.WaitGroup

type outdata struct {
	Errors []string
	Resp   []string
	Stdout string
}

var od outdata

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
	TBackup   string `json:"tbackup"`
}

func main() {

	msg, err := receive(os.Stdin)
	if err != nil {
		senderr("StdIn failed with error", err)
	}
	var data indata
	err = json.Unmarshal([]byte(msg), &data)
	if err != nil {
		senderr("UnMarshall Failed with error ", err)
	}
	sendresp("Unmarshall successful")
	if data.Path != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = ioutil.WriteFile(data.Path, []byte(data.Content), 0666)
			if err != nil {
				senderr("Save failed with error", err)
			}
			sendresp(fmt.Sprintf("Saved Successfully to %s", data.Path))
		}()
	}
	if data.Backup == "yes" {
		wg.Add(1)
		go backup(data)
	}

	wg.Wait()
	reply, _ := json.Marshal(od)
	post(reply, os.Stdout)
}

func backup(data indata) {
	defer wg.Done()
	mid := data.MessageID

	var tdir string
	if data.Bpath != "" {
		if filepath.IsAbs(data.Bpath) {
			tdir = data.Bpath
			ensuredir(tdir)
		} else {
			tdir = filepath.Join(filepath.Dir(data.Path), data.Bpath)
			ensuredir(tdir)
		}
	} else {
		tdir = filepath.Dir(data.Path)
	}
	sendresp(fmt.Sprintf("Backup path set as %s", tdir))
	var tfile = filenameWithoutExtension(filepath.Base(data.Path))
	var ext = filepath.Ext(data.Path)
	save := false
	var tfinal string
	// TOWER OF HANOI LOGIC
	if data.Bstrategy == "toh" {
		// var l, r int = data.Tohlevel, data.Tohrecent
		l, err := strconv.Atoi(data.Tohlevel)
		if err != nil {
			senderr("Error during conversion of tohlevel to int: ", err)
			return
		}
		r, err := strconv.Atoi(data.Tohrecent)
		if err != nil {
			senderr("Error during conversion of tohrecent to int: ", err)
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
				err := ioutil.WriteFile(filepath.Join(tdir, tfinal), []byte(data.Content), 0666)
				if err != nil {
					// log.Fatal(err)
					senderr("TOH Snapshot backup failed with error", err)
				}
				sendresp(fmt.Sprintf("Backed up successfully to %s", tfinal))
				save = true
				break
			}
		} // TOH SNAPSHOT LOOP ENDS

		// TOH RECENT BACKUPS
		if save != true {
			if mid < r {
				tfinal = fmt.Sprintf("%s-A%d%s", tfile, mid, ext)
				err := ioutil.WriteFile(filepath.Join(tdir, tfinal), []byte(data.Content), 0666)
				if err != nil {
					// log.Fatal(err)
					senderr("Error: TOH recent backups failed", err)
				}
				sendresp(fmt.Sprintf("Backed up successfully to %s", tfinal))
				// fmt.Printf("Saved to %s\n", tfinal)
			} else {
				c := mid % r
				tfinal = fmt.Sprintf("%s-A%d%s", tfile, c, ext)
				err := ioutil.WriteFile(filepath.Join(tdir, tfinal), []byte(data.Content), 0666)
				if err != nil {
					senderr("Error: TOH recent backups failed with error: ", err)
				}
				sendresp(fmt.Sprintf("Backed up successfully to %s", tfinal))
				// fmt.Printf("Saved to %s\n", tfinal)
			}
		} // TOH RECENT BACKUPS END
		//TOH LOGIC ENDS
	} else if data.Bstrategy == "psave" {
		pint, err := strconv.Atoi(data.Psint)
		if err != nil {
			senderr("Error during conversion of Psint to int: ", err)
			return
		}
		if mid%pint == 0 {

			tfinal = fmt.Sprintf("%s-%s%s", tfile, buildFileName(), ext)
			err := ioutil.WriteFile(filepath.Join(tdir, tfinal), []byte(data.Content), 0666)
			if err != nil {
				// log.Fatal(err)
				senderr("Error: Per save backups failed", err)
			}
			sendresp(fmt.Sprintf("Backed up successfully to %s", tfinal))
			//fmt.Printf("Saved to %s\n", tfinal)
		}
		//PERSAVE BACKUP LOGIC ENDS
	} else if data.Bstrategy == "timed" {
		if data.TBackup == "true" {
			tfinal = fmt.Sprintf("%s-%s%s", tfile, buildFileName(), ext)
			err := ioutil.WriteFile(filepath.Join(tdir, tfinal), []byte(data.Content), 0666)
			if err != nil {
				// log.Fatal(err)
				senderr("Error: Timed backups failed", err)
			}
			sendresp(fmt.Sprintf("Backed up successfully to %s", tfinal))
		} else {
			return
		}

	}
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, filepath.Ext(fn))
}

func buildFileName() string {
	return time.Now().Format("2006-01-02-15-04-05")
}

func senderr(desc string, err error) {
	apperr := fmt.Sprintf("Timimi Host: %s%s", desc, err.Error())
	od.Errors = append(od.Errors, apperr)
}
func sendresp(desc string) {
	appresp := fmt.Sprintf("Timimi Host: %s", desc)
	od.Resp = append(od.Resp, appresp)
}
func bytesToString(data []byte) string {
	return string(data[:])
}
func ensuredir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0666)
		if err != nil {
			senderr("Error: Creating backup directory failed", err)
		}
	}
}
