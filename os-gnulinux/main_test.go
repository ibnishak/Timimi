package main

import "testing"

func TestBuildfname(t *testing.T) {
	fname := buildfname("hello.html", "A2")
	if fname != "hello-A2.html" {
		t.Errorf("Sum was incorrect, got: %s, want: hello-A2.html.", fname)
	}
}

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
