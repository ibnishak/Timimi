package main

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"runtime"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type host struct {
	Exec string
}

const chrometmpl = `{
 	"name": "timimi",
  	"description": "Native messaging host to save TW5",
  	"path": "{{.Exec}}",
  	"type": "stdio",
	"allowed_origins": ["chrome-extension://mnggafnmmhdoplbffagjihajeeikgbcg/"]
}`

const firefoxtmpl = `{
	"name": "timimi",
	"description": "Native messaging host to save TW5",
	"path": "{{.Exec}}",
	"type": "stdio",
	"allowed_extensions": [ "timimi@tesseract.io" ]
  }`

func main() {
	cyan := color.New(color.FgCyan).SprintFunc() // Used when you want to mix regular output with colored output
	color.Cyan("Hello There\nStarting Timimi Installation\n\n")

	browser, err := findbrowser()
	if err != nil {
		fmt.Println("Unexpected error in choosing browser")
		os.Exit(1)
	}

	platform := runtime.GOOS
	execpath, manifestpath := findpaths(browser, platform)

	var h host
	h.Exec = execpath
	createDirIfNotExist(path.Dir(execpath))
	createDirIfNotExist(path.Dir(manifestpath))

	fmt.Println("Created host directory")

	f, err := os.Create(manifestpath) // Create host manifest file
	if err != nil {
		fmt.Println("Error while creating manifest file: ", err)
		return
	}
	defer f.Close()

	t := template.New("Timimi template") // New template
	if browser == "google-chrome" || browser == "chromium" {
		t, err = t.Parse(chrometmpl)
	} else {
		t, err = t.Parse(firefoxtmpl)
	}
	if err != nil {
		fmt.Println("Parse: ", err)
		return
	}

	err = t.Execute(f, h) // Write template "t" to file "f" with information taken from "host"
	if err != nil {
		fmt.Println("Execute: ", err)
		return
	}
	fmt.Printf("Created host manifest: %s\n", cyan(manifestpath))
	if _, err := os.Stat(execpath); os.IsNotExist(err) {
		err = os.Rename("timimi", execpath) // Rename is golang's way of moving file.
		if err != nil {
			fmt.Println("Move: ", err)
			return
		}
	}

	fmt.Printf("Created host executable: %s\n", cyan(execpath))

	color.Cyan("\n\nInstallation finished without errors.\nHave a great day!!")

}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Println("Error while creating directory", dir, err.Error())
		}
	}
}

func findbrowser() (string, error) {
	prompt := promptui.Select{
		Label: "Select browser",
		Items: []string{"firefox", "google-chrome", "chromium"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func findpaths(browser, platform string) (string, string) {
	var execpath, manifestpath string
	switch platform {
	case "darwin":
		execpath = path.Join(os.Getenv("HOME"), "Library/Application Support/timimi/timimi")
	case "linux":
		execpath = path.Join(os.Getenv("HOME"), ".local/share/timimi/timimi")
	default:
		execpath = ""
		fmt.Println("Error: Installer only intended for linux and mac. Exiting")
		os.Exit(1)
	}
	switch bp := platform + browser; bp {
	case "darwinchromium":
		manifestpath = "Library/Application Support/Chromium/NativeMessagingHosts"
	case "darwingoogle-chrome":
		manifestpath = "Library/Application Support/Google/Chrome/NativeMessagingHosts"
	case "darwinfirefox":
		manifestpath = "Library/Application Support/Mozilla/NativeMessagingHosts"
	case "linuxchromium":
		manifestpath = ".config/chromium/NativeMessagingHosts"
	case "linuxgoogle-chrome":
		manifestpath = ".config/google-chrome/NativeMessagingHosts"
	case "linuxfirefox":
		manifestpath = ".mozilla/native-messaging-hosts"
	default:
		manifestpath = ""
		fmt.Println("Error: Installer only intended for linux and mac. Exiting")
		os.Exit(1)
	}
	manifestpath = path.Join(os.Getenv("HOME"), manifestpath, "timimi.json")
	return execpath, manifestpath
}
