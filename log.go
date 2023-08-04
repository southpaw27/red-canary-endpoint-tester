package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Log struct {
	ProcessLog []ProcessLog `json:"processLog"`
	FileLog    []FileLog    `json:"fileLog"`
	NetworkLog []NetworkLog `json:"networkLog"`
}

type ProcessLog struct {
	Timestamp time.Time `json:"timestamp"`
	Profile   Profile   `json:"profile"`
}

type FileLog struct {
	FilePath   string     `json:"filepath"`
	Actions    []string   `json:"actions"`
	ProcessLog ProcessLog `json:"processLog"`
}

type NetworkLog struct {
	Source      string     `json:"source"`
	Destination string     `json:"destination"`
	DataAmount  int64      `json:"dataAmount"`
	Protocol    string     `json:"protocol"`
	ProcessLog  ProcessLog `json:"processLog"`
}

// Log process level information
func logProcess(p *Profile) {
	log := processLog(p)

	globalLog.ProcessLog = append(globalLog.ProcessLog, log)
}

// Create a process log
func processLog(p *Profile) ProcessLog {
	if p == nil {
		return ProcessLog{
			Timestamp: time.Now(),
			Profile:   globalProfile,
		}
	} else {
		return ProcessLog{
			Timestamp: time.Now(),
			Profile:   *p,
		}
	}
}

// Create a file activity log
func logFileActivity(filepath string, actions []string) {
	log := FileLog{
		FilePath:   filepath,
		Actions:    actions,
		ProcessLog: processLog(nil),
	}

	globalLog.FileLog = append(globalLog.FileLog, log)
}

// Create a network activity log
func logNetworkActivity(source string, destination string, dataAmount int64, protocol string) {
	log := NetworkLog{
		Source:      source,
		Destination: destination,
		DataAmount:  dataAmount,
		Protocol:    protocol,
		ProcessLog:  processLog(nil),
	}

	globalLog.NetworkLog = append(globalLog.NetworkLog, log)
}

// Print the logs to file
func logsToFile(args Arguments) {
	jsonData, err := json.MarshalIndent(globalLog, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling global log to JSON", err)
		return
	}

	file, err := os.Create(args.logFile)

	if err != nil {
		fmt.Println("Error creating log file", err)
		return
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to log file", err)
		return
	}
}
