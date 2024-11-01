package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

//This  function will retrieve the width and the height of the terminal.
func TerminalDimensions() (width, height int) {
	var winsize struct {
		Row uint16
		Col uint16
		Xpx uint16
		Ypx uint16
	}

	fd := int(os.Stdout.Fd())

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&winsize)))
	if err != 0 {
		return 0, 0
	}

	return int(winsize.Col), int(winsize.Row)
}
