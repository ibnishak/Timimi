package main

import "testing"

func TestSetbackuppath(t *testing.T) {
	tables := []struct {
		bpath, path, finalpath string
	}{
		{"", "/a/b/c.html", "/a/b"},
		{"/a/b/c", "/a/b/c.html", "/a/b/c"},
		{"..", "/a/b/c.html", "/a"},
	}

	for _, table := range tables {
		result := setbackuppath(table.bpath, table.path)
		if result != table.finalpath {
			t.Errorf("For paths %s and %s, Wanted %s, Got: %s", table.bpath, table.path, table.finalpath, result)
		}
	}
}

func TestFilenameWithoutExtension(t *testing.T) {
	tables := []struct {
		a, b string
	}{
		{"/a/b/c.html", "/a/b/c"},
		{"/a/b/c", "/a/b/c"},
	}

	for _, table := range tables {
		result := filenameWithoutExtension(table.a)
		if result != table.b {
			t.Errorf("For paths %s, Wanted %s, Got: %s", table.a, table.b, result)
		}
	}
}
