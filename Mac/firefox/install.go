package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"

	"github.com/fatih/color"
)

type Host struct {
	Dir      string
	Manifest string
	Exec     string
}

// Template for timimi host manifest
const tmpl = `{
  "name": "timimi",
  "description": "Native messaging host to save TW5",
  "path": "{{.Exec}}",
  "type": "stdio",
  "allowed_extensions": [ "timimi@tesseract.com" ]
}
`

func main() {
	color.Cyan("Hello There\nStarting Timimi Installation")
	cyan := color.New(color.FgCyan).SprintFunc() // Used when you want to mix regular output with colored output

	var host Host
	host.Dir = path.Join(os.Getenv("HOME"), "Library/Application Support/Mozilla/NativeMessagingHosts")
	host.Manifest = path.Join(host.Dir, "timimi.json")
	host.Exec = path.Join(host.Dir, "timimi")

	createDirIfNotExist(host.Dir) // create native host directory
	fmt.Printf("Created host directory: %s\n", cyan(host.Dir))

	f, err := os.Create(host.Manifest) // Create host manifest file
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

	err = t.Execute(f, host) // Write template "t" to file "f" with information taken from "host"
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
	fmt.Printf("Created host manifest: %s\n", cyan(host.Manifest))

	err = os.Rename("timimi", host.Exec) // Rename is golang's way of moving file.
	if err != nil {
		log.Fatal("Move: ", err)
		return
	}
	fmt.Printf("Created host executable: %s\n", cyan(host.Exec))

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
