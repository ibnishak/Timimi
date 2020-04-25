package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/beeep"
)

func ensuredir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0777)
	}
	return nil
}

func notifier() {
	beeep.Notify("Error!", "Error while saving data by Timimi Host. Please check error.log for details", "")
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, filepath.Ext(fn))
}

func setbackuppath(bpath, path string) string {
	if bpath == "" {
		return filepath.Dir(path)
	} else if filepath.IsAbs(bpath) {
		return bpath
	} else {
		return filepath.Join(filepath.Dir(path), bpath)
	}
}

func unmarshdata(msg []byte) (indata, error) {
	var data indata
	err := json.Unmarshal(msg, &data)
	if err != nil || data.Path == "" {
		return indata{}, err
	}
	return data, nil
}

func logerr(a string, b string) {
	f, err := os.OpenFile("timimi.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "Timimi", log.LstdFlags)
	logger.Println(a, b)
}
