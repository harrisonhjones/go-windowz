// +build windows

package windowz

import (
	"syscall"
)

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	procSetWindowPos = user32.MustFindProc("SetWindowPos")
)

type ZOrder uint

// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms633545(v=vs.85).aspx
const (
	WINDOW_BOTTOM    ZOrder = 1
	WINDOW_NOTOPMOST ZOrder = ZOrder(^uint(0) - 1)
	WINDOW_TOP       ZOrder = 0
	WINDOW_TOPMOST   ZOrder = ZOrder(^uint(0))
)

// SetWindowZ sets a window's Z order.
// An error is returned if encountered.
func Set(hwnd syscall.Handle, zOrder ZOrder) error {
	ret, _, err := procSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(zOrder),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0x0001|0x0002))

	if ret == 0 {
		return err
	}
	return nil
}
