package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"time"
)

func timed(tbackup, path string) (string, bool) {
	if tbackup == "true" {
		return buildfname(path, ""), true
	}
	return "", false
}

func psave(psint, path string, mid int) (string, bool) {
	// TODO: Ensure that psint is an integer
	pint, _ := strconv.Atoi(psint)
	if mid%pint == 0 {
		return buildfname(path, ""), true
	}
	return "", false
}

func toh(path string, l, r, mid int) (string, bool) {
	for n := float64(l); n >= 3; n-- {
		if mid%8 != 0 {
			break
		}
		p := int(math.Pow(2, n))
		if mid%p == 0 {
			return buildfname(path, fmt.Sprintf("%d", p)), true
		}
	}

	if mid < r {
		return buildfname(path, fmt.Sprintf("A%d", mid)), true
	}

	c := mid % r
	return buildfname(path, fmt.Sprintf("A%d", c)), true

}

func buildfname(path, uniq string) string {
	title := filenameWithoutExtension(filepath.Base(path))
	ext := filepath.Ext(path)
	if uniq == "" {
		uniq = time.Now().Format("2006-01-02-15-04-05")
	}
	return fmt.Sprintf("%s-%s%s", title, uniq, ext)
}

func fifo(fifoint string, mid int, path string) (string, bool) {
	fint, _ := strconv.Atoi(fifoint)
	b := mid % fint
	return buildfname(path, fmt.Sprintf("Backup-%d", b)), true
}
