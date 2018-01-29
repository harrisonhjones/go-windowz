// +build windows

package main

import (
	"flag"
	"fmt"
	"syscall"
	"time"

	"harrisonhjones.com/go-window-z"
)

var (
	user32 = syscall.MustLoadDLL("user32.dll")

	procGetForegroundWindow = user32.MustFindProc("GetForegroundWindow")
	procGetWindowTextW      = user32.MustFindProc("GetWindowTextW")
	procBeep                = user32.MustFindProc("MessageBeep")

	flagDuration string
	flagZ        string
)

func init() {
	flag.StringVar(&flagDuration, "d", "5s", "the amount of time to wait before setting the active window's Z order")
	flag.StringVar(&flagZ, "z", "topmost", "the z order to set the foreground window to [topmost, top, nottopmost, bot")
}

func main() {
	fmt.Printf("windowz\n\tWritten by Harrison Jones (harrison@hhj.me)\n\tSets the z order of the foreground window after a delay\n\n")
	flag.Parse()

	duration := mustParseSleepDuration(flagDuration)
	zOrder := mustParseZOrder(flagZ)

	procBeep.Call()
	time.Sleep(duration)
	procBeep.Call()

	h := mustGetForegroundWindow()
	windowTitle := mustGetWindowText(h)
	if err := windowz.Set(h, zOrder); err != nil {
		fatal("Failed to set the z order: %+v", err)
	}

	fmt.Printf("Window (%s) set to %s", windowTitle, flagZ)
}
