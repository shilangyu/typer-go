package utils

import (
	"net"
	"strings"
)

// StringDimensions returns the width and height of a string
// where w = longest line, h = amount of lines
func StringDimensions(s string) (w, h int) {
	text := strings.Split(s, "\n")
	h = len(text)

	for _, line := range text {
		if len(line) > w {
			w = len(line)
		}
	}

	return
}

// Center takes an array of strings and adds spaces to center them
func Center(s []string) (res []string) {
	maxW, _ := StringDimensions(strings.Join(s, "\n"))

	for _, text := range s {
		diff := maxW - len(text)
		text = strings.Repeat(" ", diff/2) + text + strings.Repeat(" ", diff-diff/2)
		res = append(res, text)
	}

	return
}

// IPv4 returns users ipv4 as a string
func IPv4() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	myIP := localAddr.IP.String()
	conn.Close()
	return myIP
}
