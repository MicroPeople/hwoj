package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

// var commands = map[string]string{
// 	"windows": "cmd",
// 	"darwin":  "open",
// 	"linux":   "xdg-open",
// }

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	fmt.Println(args)
	return exec.Command(cmd, args...).Start()
}
