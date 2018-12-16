package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
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
}
type outdata struct {
}

func main() {

	msg, err := receive(os.Stdin)
	if err != nil {
		panic(err)
	}
	var data indata
	err = json.Unmarshal([]byte(msg), &data)
	if err != nil {
		send("UnMarshall Failed")
		panic(err)
	}
	send("UnMarshall Sucess")
	if data.Path != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			send("Entering Save branch")
			err = ioutil.WriteFile(data.Path, []byte(data.Content), 0666)
			if err != nil {
				log.Fatal(err)
				send("Save Failed")
			}
			send("Save Successfull")
		}()
	}
	if data.Backup == "yes" {
		send("Entered Backup")
		wg.Add(1)
		go backup(data)
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
			send("Error during conversion of tohlevel to int")
			log.Fatal("Error during conversion of tohlevel to int: ", err)
			return
		}
		r, err := strconv.Atoi(data.Tohrecent)
		if err != nil {
			send("Error during conversion of tohrecent to int")
			log.Fatal("Error during conversion of tohrecent to int: ", err)
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
					log.Fatal(err)
				}
				fmt.Printf("Saved to %s\n", tfinal)
				send("Backup Success")
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
					log.Fatal(err)
				}

				fmt.Printf("Saved to %s\n", tfinal)
				send("Backup Success")

			} else {
				c := mid % r
				tfinal = fmt.Sprintf("%s-A%d%s", tfile, c, ext)
				err := ioutil.WriteFile(path.Join(tdir, tfinal), []byte(data.Content), 0666)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Saved to %s\n", tfinal)
				send("Backup Success")
			}
		} // TOH RECENT BACKUPS END
		//TOH LOGIC ENDS
	} else if data.Bstrategy == "psave" {
		pint, err := strconv.Atoi(data.Psint)
		if err != nil {
			log.Fatal("Error during conversion of Psint to int: ", err)
			return
		}
		if mid%pint == 0 {

			tfinal = fmt.Sprintf("%s-%s%s", tfile, buildFileName(), ext)
			err := ioutil.WriteFile(path.Join(tdir, tfinal), []byte(data.Content), 0666)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Saved to %s\n", tfinal)
		}
	} //PERSAVE BACKUP LOGIC ENDS
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func buildFileName() string {
	return time.Now().Format("2006-01-02-15-04-05")
}

func send(res string) {
	reply := fmt.Sprintf(`{ "content":"%s"}`, res)
	post([]byte(reply), os.Stdout)
}
