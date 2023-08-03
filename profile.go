package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

type Profile struct {
	CurrentUser        string   `json:"currentUser"`
	OperatingSystem    string   `json:"operatingSystem"`
	ProcessName        string   `json:"processName"`
	ProcessCommandLine []string `json:"processCommandLine"`
	ProcessId          int      `json:"processId"`
}

func (p *Profile) init() {
	p.setCurrentUser()
	p.setOperatingSystem()
	p.setProcessId()
	p.setProcessCommandLine()
	p.setProcessName()
}

// return the current user
func currentUser() string {
	currentUser, err := user.Current()

	if err != nil {
		fmt.Println("Failed to get the current user:", err)
		return ""
	}

	return currentUser.Username
}

func (p *Profile) setCurrentUser() {
	p.CurrentUser = currentUser()
}

func (p *Profile) setOperatingSystem() {
	p.OperatingSystem = operatingSystem()
}

func operatingSystem() string {
	return runtime.GOOS
}

func processId() int {
	return os.Getpid()
}

func (p *Profile) setProcessId() {
	p.ProcessId = processId()
}

func processCommandLine() []string {
	return os.Args
}

func (p *Profile) setProcessCommandLine() {
	p.ProcessCommandLine = processCommandLine()
}

func processName() string {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error retrieving executable path:", err)
		return ""
	}
	return filepath.Base(execPath)
}

func (p *Profile) setProcessName() {
	p.ProcessName = processName()
}
