package main

import (
	"syscall"
	"unsafe"
)

const (
	DRIVE_TYPE_UNKNOWN    = iota
	DRIVE_TYPE_WRONG_PATH = iota
	DRIVE_TYPE_REMOVABLE  = iota
	DRIVE_TYPE_FIXED      = iota
	DRIVE_TYPE_REMOTE     = iota
	DRIVE_TYPE_CDROM      = iota
	DRIVE_TYPE_RAMDISK    = iota
)

func GetDriveType(drivePath string) int {
	utf16DrivePath, err := syscall.UTF16FromString(drivePath)
	if err != nil {
		panic(err)
	}
	procGetDriveTypeW := syscall.NewLazyDLL("kernel32.dll").
		NewProc("GetDriveTypeW")
	drive_type, _, _ := syscall.Syscall(
		procGetDriveTypeW.Addr(),
		1,
		uintptr(unsafe.Pointer(&utf16DrivePath[0])),
		0,
		0,
	)
	return int(drive_type)
}
