// +build windows

package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"harrisonhjones.com/go-windowz"
)

func mustGetForegroundWindow() syscall.Handle {
	r0, _, err := procGetForegroundWindow.Call()
	if r0 == 0 {
		fatal("Failed to get the foreground window: %+v", err)
		return 0
	}
	return syscall.Handle(r0)
}

func mustGetWindowText(hwnd syscall.Handle) string {
	b := make([]uint16, 200)
	r0, _, err := procGetWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&b[0])), uintptr(int32(len(b))))
	if int32(r0) == 0 {
		fatal("Failed to get the window title: %+v", err)
		return ""
	}
	return syscall.UTF16ToString(b)
}

func mustParseSleepDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		fatal("Failed to parse specified sleep duration: %+v", err)
	}
	return d
}

func mustParseZOrder(s string) windowz.ZOrder {
	switch strings.ToLower(s) {
	case "topmost":
		return windowz.WINDOW_TOPMOST
	case "top":
		return windowz.WINDOW_TOP
	case "nottopmost":
		return windowz.WINDOW_NOTOPMOST
	case "bot":
		return windowz.WINDOW_BOTTOM
	default:
		fatal("%s is not a supported z order", s)
		return 0
	}
}

func fatal(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}
