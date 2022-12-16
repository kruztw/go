package utils

import (
	"bytes"
	"encoding/base64"
	"os/exec"
	"simple/logger"
	"strings"
	"text/template"
	"unsafe"

	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/unicode"
)

type DWORD uint32
type UCHAR byte

type WTS_SESSION_INFOA struct {
	SessionId      DWORD
	WinStationName *UCHAR
	State          DWORD
}

const (
	WTSUserName = 5
	username    = "raybai"
)

var (
	WTS_CURRENT_SERVER_HANDLE       = uintptr(0)
	modWtsapi32                     = windows.NewLazyDLL("wtsapi32.dll")
	procWTSQuerySessionInformationA = modWtsapi32.NewProc("WTSQuerySessionInformationA")
	procWTSQueryUserToken           = modWtsapi32.NewProc("WTSQueryUserToken")
	procWTSEnumerateSessionsA       = modWtsapi32.NewProc("WTSEnumerateSessionsA")
	procWTSFreeMemory               = modWtsapi32.NewProc("WTSFreeMemory")
	toastTemplate                   *template.Template
)

func init() {
	toastTemplate = template.New("toast")
	_, _ = toastTemplate.Parse(`get-credential username`)
}

func runAsUser(token windows.Token, program, cwd string, args []string) (err error) {
	b := new(strings.Builder)
	b.WriteString(program)
	for _, a := range args {
		b.WriteByte(' ')
		b.WriteString(a)
	}

	command, err := windows.UTF16PtrFromString(b.String())
	if err != nil {
		logger.Log.Errorf("cannot encode command to UTF16: %v", err)
		return
	}

	programPtr, err := windows.UTF16PtrFromString(program)
	if err != nil {
		logger.Log.Errorf("cannot encode program to UTF16: %v", err)
		return
	}

	cwdPtr, err := windows.UTF16PtrFromString(cwd)
	if err != nil {
		logger.Log.Errorf("cannot encode cwd to UTF16: %v", err)
		return
	}

	startupinfo := windows.StartupInfo{
		Flags:      windows.STARTF_USESTDHANDLES | windows.STARTF_USESHOWWINDOW,
		ShowWindow: windows.SW_SHOW,
	}
	procinfo := windows.ProcessInformation{
		Process: windows.InvalidHandle,
		Thread:  windows.InvalidHandle,
	}

	err = windows.CreateProcessAsUser(
		token,
		programPtr,
		command,
		nil,
		nil,
		false,
		0,
		nil,    // environment
		cwdPtr, // current directory
		&startupinfo,
		&procinfo,
	)

	if err != nil {
		logger.Log.Errorf("CreateProcessAsUser: %v", err)
		return err
	}

	if rc, lastErr := windows.WaitForSingleObject(procinfo.Process, windows.INFINITE); rc != 0 {
		logger.Log.Errorf("WaitForSingleObject, rc=%v, err=%v", rc, lastErr)
		return lastErr
	}

	return
}

func toastNotify(token windows.Token) (err error) {
	encodedCommand, err := buildEncodedCommand(toastTemplate)
	if err != nil {
		return
	}

	execPath, err := exec.LookPath("powershell.exe")
	if err != nil {
		logger.Log.Error(err)
		return
	}

	return runAsUser(token, execPath, ".", []string{
		"-ExecutionPolicy", "RemoteSigned",
		"-EncodedCommand", string(encodedCommand),
	})
}

func buildEncodedCommand(tmpl *template.Template) ([]byte, error) {
	var out bytes.Buffer
	b64Encoder := base64.NewEncoder(base64.StdEncoding, &out)
	utf16Encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	utf16Writer := utf16Encoder.Writer(b64Encoder)

	if err := tmpl.Execute(utf16Writer, nil); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func GetCredential() {
	sessionIds, err := getLogonSessionIds(false)
	if err != nil {
		logger.Log.Errorf("get session ids failed: %v", err)
		return
	}

	if len(sessionIds) < 1 {
		logger.Log.Errorf("sessionIds < 1")
		return
	}

	logger.Log.Infof("sessionIds = %v", sessionIds)

	for _, sessionId := range sessionIds {
		go func(sessionId int) {
			if token, err := duplicateUserToken(sessionId); err != nil {
				logger.Log.Warnf("get user token failed: %v", err)
			} else {
				if err := toastNotify(token); err != nil {
					logger.Log.Warnf("send notification failed: %v", err)
				}
			}
		}(sessionId)
	}
}

func getLogonSessionIds(force bool) ([]int, error) {
	var err error

	sessionIds, err := FindLoginUserSessionIds(username)
	if err != nil {
		return nil, err
	}

	return sessionIds, nil
}

func FindLoginUserSessionIds(username string) (sessionIds []int, err error) {
	var reserved uint32
	var count uint32
	var sessions *WTS_SESSION_INFOA

	sessionIds = []int{}

	success, _, lastErr := procWTSEnumerateSessionsA.Call(
		WTS_CURRENT_SERVER_HANDLE,
		uintptr(reserved),
		uintptr(1),
		uintptr(unsafe.Pointer(&sessions)),
		uintptr(unsafe.Pointer(&count)),
	)

	logger.Log.Infof("success = %v, lastErr = %v, sessions = %v", success, lastErr, sessions)

	defer func() {
		if sessions != nil {
			_, _, _ = procWTSFreeMemory.Call(
				uintptr(unsafe.Pointer(sessions)),
			)
		}
	}()

	if success == 0 {
		return nil, lastErr
	}

	for _, user := range unsafe.Slice(sessions, count) {
		name, err := getUsername(user.SessionId)
		if err != nil {
			return nil, err
		} else if strings.EqualFold(username, name) {
			sessionIds = append(sessionIds, int(user.SessionId))
		}
	}

	return
}

func getUsername(sessionId DWORD) (username string, err error) {
	var namePtr *byte
	var bytesReturned DWORD

	success, _, lastErr := procWTSQuerySessionInformationA.Call(
		WTS_CURRENT_SERVER_HANDLE,
		uintptr(sessionId),
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
		return "", lastErr
	}

	if bytesReturned == 0 || *namePtr == 0 {
		return "", nil
	}

	return string(unsafe.Slice(namePtr, bytesReturned)[:bytesReturned-1]), nil
}

func duplicateUserToken(sessionID int) (windows.Token, error) {
	var token windows.Token
	success, _, lastErr := procWTSQueryUserToken.Call(
		uintptr(sessionID),
		uintptr(unsafe.Pointer(&token)),
	)

	if success == 0 {
		return windows.Token(0), lastErr
	}

	var duplicatedToken windows.Token
	err := windows.DuplicateTokenEx(
		token,
		windows.MAXIMUM_ALLOWED,
		nil,
		windows.SecurityDelegation,
		windows.TokenPrimary,
		&duplicatedToken,
	)
	if err != nil {
		return 0, err
	}

	return duplicatedToken, nil
}
