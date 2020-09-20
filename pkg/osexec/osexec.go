package osexec

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/ge"
	osDetector "github.com/nobuenhombre/suikat/pkg/os-detector"
)

type RunError struct {
	Command string
	Args    []string
	Parent  error
	StdOut  string
}

func (e *RunError) Error() string {
	return fmt.Sprintf(
		"exec error\n{\tCommand:\t%v,\nArgs:\t\t%v,\nParentError:\t%v,\nStdOut:\t%v}",
		e.Command,
		strings.Join(e.Args, ",\n\t\t"),
		e.Parent,
		e.StdOut,
	)
}

func Run(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), &RunError{
			Command: command,
			Args:    args,
			Parent:  err,
			StdOut:  string(out),
		}
	}

	return string(out), nil
}

func OSRun(command string, args []string) (string, error) {
	switch runtime.GOOS {
	case osDetector.OSLinux, osDetector.OSMacOs:
		return Run(command, args)

	case osDetector.OSWindows:
		return "", &ge.NotReleasedError{
			Name: "NotReleasedError for OS Windows",
		}

	default:
		return "", &osDetector.UnknownOSError{
			Name: runtime.GOOS,
		}
	}
}
