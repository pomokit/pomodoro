package main

import (
	"syscall"
	"unsafe"
)

func setConsoleTitle(title string) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTitleW")
	proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
}
