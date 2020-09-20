package osdetector

import (
	"fmt"
)

const (
	OSWindows = "windows"
	OSLinux   = "linux"
	OSMacOs   = "darwin"
	OSUnknown = "unknown"
)

type UnknownOSError struct {
	Name string
}

func (e *UnknownOSError) Error() string {
	return fmt.Sprintf("Unknown OS (name = %v)", e.Name)
}
