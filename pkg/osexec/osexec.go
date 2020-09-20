package osexec

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

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

func RunUnix(command string, args []string) (string, error) {
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

func RunWindows(command string, args []string) (string, error) {
	arg := fmt.Sprintf("%v %v", command, strings.Join(args, " "))
	cmd := exec.Command("CMD", "/c", arg)

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
		return RunUnix(command, args)

	case osDetector.OSWindows:
		return RunWindows(command, args)

	default:
		return "", &osDetector.UnknownOSError{
			Name: runtime.GOOS,
		}
	}
}
