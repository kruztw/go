package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	// https://learn.microsoft.com/en-us/windows/win32/api/wtsapi32/ne-wtsapi32-wts_info_class
	WTSUserName = 5
)

var (
	WTS_CURRENT_SERVER_HANDLE       = uintptr(0)
	modWtsapi32                     = windows.NewLazyDLL("wtsapi32.dll")
	procWTSQuerySessionInformationA = modWtsapi32.NewProc("WTSQuerySessionInformationA")
	procWTSEnumerateSessionsA       = modWtsapi32.NewProc("WTSEnumerateSessionsA")
	procWTSFreeMemory               = modWtsapi32.NewProc("WTSFreeMemory")
)

type WTS_SESSION_INFOA struct {
	SessionId      uint32
	WinStationName *byte
	State          uint32
}

func main() {
	var reserved uint32
	var count uint32
	var sessions *WTS_SESSION_INFOA

	success, _, lastErr := procWTSEnumerateSessionsA.Call(
		WTS_CURRENT_SERVER_HANDLE,
		uintptr(reserved),
		uintptr(1),
		uintptr(unsafe.Pointer(&sessions)),
		uintptr(unsafe.Pointer(&count)),
	)

	if success == 0 {
		fmt.Printf("procWTSEnumerateSessionsA failed: %v", lastErr)
		return
	}

	for _, session := range unsafe.Slice(sessions, count) {
		fmt.Printf("session = %v\n", session)
		var namePtr *byte
		var bytesReturned uint32

		success, _, lastErr := procWTSQuerySessionInformationA.Call(
			WTS_CURRENT_SERVER_HANDLE,
			uintptr(session.SessionId),
			uintptr(WTSUserName),
			uintptr(unsafe.Pointer(&namePtr)),
			uintptr(unsafe.Pointer(&bytesReturned)),
		)

		defer func() {
			if namePtr != nil {
				_, _, _ = procWTSFreeMemory.Call(
					uintptr(unsafe.Pointer(namePtr)),
				)
			}
		}()

		if success == 0 {
			fmt.Printf("procWTSQuerySessionInformationA failed: %v", lastErr)
			return
		}

		fmt.Printf("username = %v\n", namePtr)
	}
}
