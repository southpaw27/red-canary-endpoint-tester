package main

import "flag"

// Hold all of our arguments, referenced throughout
type Arguments struct {
	consolePrint  string
	create        bool
	delete        bool
	endpoint      string
	execFilePath  string
	execArgs      string
	filepath      string
	fileString    string
	logFile       string
	modify        bool
	modifyAdd     bool
	modifyReplace bool
	opts          string
	sendFile      bool
}

// Define and parse the flags into a centralized struct
func collectFlags(args *Arguments) {
	flag.StringVar(&args.endpoint, "ep", "", "The endpoint to connect and send data to")
	flag.StringVar(&args.execArgs, "ea", "", "Arguments to run the exec in '-arg1=123|-arg2=456' form separated by a pipe")
	flag.StringVar(&args.execFilePath, "ef", "", "exec file path to run")
	flag.StringVar(&args.filepath, "fp", "", "The filepath")
	flag.StringVar(&args.fileString, "fs", "", "The string to add to a file")
	flag.StringVar(&args.logFile, "lf", "./logFile.json", "The file path to send logs to")

	flag.BoolVar(&args.create, "c", false, "Create a new file at filepath -fp")
	flag.BoolVar(&args.delete, "d", false, "Delete a file at filepath -fp")
	flag.BoolVar(&args.modify, "m", false, "Modify a file at filepath -fp")
	flag.BoolVar(&args.modifyAdd, "ma", false, "Modify a file at filepath -fp by adding to the end of the file")
	flag.BoolVar(&args.modifyAdd, "mr", false, "Modify a file at filepath -fp by replacing the contents of the file")
	flag.BoolVar(&args.sendFile, "sf", false, "If there is a file at -fp, and this flag is true, we will send the file to the endpoint specified in -ep")

	// Parse the flags from the command line
	flag.Parse()
}
