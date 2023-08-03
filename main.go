package main

/*
	What this program can do:
		x Start a process, given a path to an exec file and the desired command-line arugments
		x Create a file of a specified type at a specified location
		x Modify a file
		x Delete a file
		x Establish a network connection and transmit data
*/

// Have a global log for easy logging anywhere.
var globalLog = Log{
	ProcessLog: make([]ProcessLog, 0),
	FileLog:    make([]FileLog, 0),
	NetworkLog: make([]NetworkLog, 0),
}

var globalProfile = Profile{}

func main() {
	var args Arguments

	// This is the profile for the main process. It's not going to change midway, and we use it
	// quite a bit in logging, so just make a global one to reuse
	globalProfile.init()

	// Parse the flags from the command line
	collectFlags(&args)

	// Log the current process
	logProcess(nil)
	// Start any processes passed from the command line
	startProcess(args)
	// Create/Update/Delete any files from the command line
	fileAction(args)
	// Perform network actions from the command line
	connectAndSendData(args)
	// Print the logs to the log file
	logsToFile(args)
}
