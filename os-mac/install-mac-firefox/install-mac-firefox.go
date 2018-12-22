package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"

	"github.com/fatih/color"
)

type host struct {
	Dir      string
	Manifest string
	Exec     string
}

var pldir = "Library/Application Support/Mozilla/NativeMessagingHosts" //Platform Directory

// Template for timimi host manifest
const tmpl = `{
  "name": "timimi",
  "description": "Native messaging host to save TW5",
  "path": "{{.Exec}}",
  "type": "stdio",
  "allowed_extensions": [ "timimi@tesseract.io" ]
}
`

func main() {
	color.Cyan("Hello There\nStarting Timimi Installation")
	cyan := color.New(color.FgCyan).SprintFunc() // Used when you want to mix regular output with colored output

	var h host
	h.Dir = path.Join(os.Getenv("HOME"), pldir)
	h.Manifest = path.Join(h.Dir, "timimi.json")
	h.Exec = path.Join(h.Dir, "timimi")

	createDirIfNotExist(h.Dir) // create native host directory
	fmt.Printf("Created host directory: %s\n", cyan(h.Dir))

	f, err := os.Create(h.Manifest) // Create host manifest file
	if err != nil {
		log.Fatal("Create file: ", err)
		return
	}
	defer f.Close()

	t := template.New("Timimi template") // New template
	t, err = t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err = t.Execute(f, h) // Write template "t" to file "f" with information taken from "host"
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
	fmt.Printf("Created host manifest: %s\n", cyan(h.Manifest))

	err = os.Rename("timimi", h.Exec) // Rename is golang's way of moving file.
	if err != nil {
		log.Fatal("Move: ", err)
		return
	}
	fmt.Printf("Created host executable: %s\n", cyan(h.Exec))

	color.Cyan("\n\nInstallation finished without errors.\nHave a great day!!")

}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
