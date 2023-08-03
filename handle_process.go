package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// Start any process that is passed in via the command line
func startProcess(args Arguments) {
	if args.execFilePath == "" {
		return
	}
	// Get the absolute file path
	var cleanedExecPath, err = filepath.Abs(filepath.Join(filepath.Split(args.execFilePath)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var processName = filepath.Base(cleanedExecPath)

	// Path to the exec
	// Arguments
	opts := strings.Split(args.execArgs, "|")

	// Create the command to run
	cmd := exec.Command(cleanedExecPath, opts...)

	// Run the command
	err = cmd.Start()

	processProfile := Profile{
		CurrentUser:        globalProfile.CurrentUser,
		OperatingSystem:    globalProfile.OperatingSystem,
		ProcessName:        processName,
		ProcessCommandLine: cmd.Args,
		ProcessId:          cmd.Process.Pid,
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logProcess(&processProfile)
}
