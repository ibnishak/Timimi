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

func notifier(msg string, logger *log.Logger) {
	logger.Println("Error while saving data from browser.", msg)
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
