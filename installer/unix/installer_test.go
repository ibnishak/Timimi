package main

import "testing"

func TestFindpaths(t *testing.T) {

	tables := []struct {
		browser      string
		platform     string
		execpath     string
		manifestpath string
	}{
		{"chromium", "darwin", "/home/richie/Library/Application Support/timimi/timimi", "/home/richie/Library/Application Support/Chromium/NativeMessagingHosts/timimi.json"},
		{"google-chrome", "darwin", "/home/richie/Library/Application Support/timimi/timimi", "/home/richie/Library/Application Support/Google/Chrome/NativeMessagingHosts/timimi.json"},
		{"firefox", "darwin", "/home/richie/Library/Application Support/timimi/timimi", "/home/richie/Library/Application Support/Mozilla/NativeMessagingHosts/timimi.json"},
		{"chromium", "linux", "/home/richie/.local/share/timimi/timimi", "/home/richie/.config/chromium/NativeMessagingHosts/timimi.json"},
		{"google-chrome", "linux", "/home/richie/.local/share/timimi/timimi", "/home/richie/.config/google-chrome/NativeMessagingHosts/timimi.json"},
		{"firefox", "linux", "/home/richie/.local/share/timimi/timimi", "/home/richie/.mozilla/native-messaging-hosts/timimi.json"},
	}

	for _, table := range tables {
		e, m := findpaths(table.browser, table.platform)
		if e != table.execpath || m != table.manifestpath {
			t.Errorf("For %s and %s, we wanted %s and %s \n but received %s and %s", table.browser, table.platform, table.execpath, table.manifestpath, e, m)
		}
	}
}
