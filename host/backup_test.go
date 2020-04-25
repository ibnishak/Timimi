package main

import "testing"

func TestToh(t *testing.T) {
	tables := []struct {
		l       int
		r       int
		mid     int
		bname   string
		bstatus bool
	}{
		{3, 5, 4, "test-A4.html", true},
		{3, 5, 6, "test-A1.html", true},
		{3, 5, 8, "test-8.html", true},
		{4, 5, 16, "test-16.html", true},
		{4, 5, 24, "test-8.html", true},
		{4, 5, 23, "test-A3.html", true},
	}

	for _, table := range tables {
		n, s := toh("/home/richie/Repos/test/test.html", table.l, table.r, table.mid)
		if n != table.bname || s != table.bstatus {
			t.Errorf("\nError: Want %s and %t\n Got %s and %t", table.bname, table.bstatus, n, s)
		}
	}
}

func TestTimed(t *testing.T) {
	tables := []struct {
		tbackup string
		bstatus bool
	}{
		{"true", true},
		{"false", false},
	}
	for _, table := range tables {
		_, s := timed(table.tbackup, "/home/richie/Repos/test/test.html")
		if s != table.bstatus {
			t.Errorf("\nError: For %s, want %t, got %t", table.tbackup, table.bstatus, s)
		}
	}
}

func TestPsave(t *testing.T) {
	tables := []struct {
		i string
		m int
		s bool
	}{
		{"4", 8, true},
		{"4", 7, false},
	}
	for _, table := range tables {
		_, result := psave(table.i, "/home/richie/Repos/test/test.html", table.m)
		if result != table.s {
			t.Errorf("Error: For %s and %d, \nWANT: %t\nGOT: %t", table.i, table.m, table.s, result)
		}
	}
}

func TestBuildfname(t *testing.T) {
	fname := buildfname("hello.html", "A2")
	if fname != "hello-A2.html" {
		t.Errorf("Error, got: %s, want: hello-A2.html.", fname)
	}
}

func TestFifo(t *testing.T) {
	tables := []struct {
		i string
		m int
		n string
	}{
		{"5", 5, "hello-Backup-0.html"},
		{"5", 6, "hello-Backup-1.html"},
		{"5", 3, "hello-Backup-3.html"},
	}
	for _, table := range tables {
		result, _ := fifo(table.i, table.m, "/a/b/c/hello.html")
		if result != table.n {
			t.Errorf("\nError: For %s and %d\nWANT: %s\nGOT: %s", table.i, table.m, table.n, result)
		}
	}
}
