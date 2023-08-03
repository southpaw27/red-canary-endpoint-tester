package main

import (
	"errors"
	"fmt"
	"os"
)

// Handle file actions that have been passed in
func fileAction(args Arguments) {
	// See if we have any actions to take, exit early if not
	if !(args.create || args.modify || args.delete || args.modifyAdd || args.modifyReplace) {
		return
	}

	if args.delete {
		err := os.Remove(args.filepath)

		if err != nil {
			fmt.Println("Error deleting the file")
		}
		logFileActivity(args.filepath, actionMap(args))
		return
	}

	file, err := openFile(args)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}

	// make sure the file will close after we return from the method
	defer file.Close()

	if args.fileString != "" {
		_, err = file.WriteString(args.fileString)

		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	logFileActivity(args.filepath, actionMap(args))
}

// Open the file in the correct mode for the passed in flags
func openFile(args Arguments) (*os.File, error) {
	if args.filepath == "" {
		return nil, errors.New("No filepath given")
	}
	if args.create || args.modifyReplace {
		return os.Create(args.filepath)
	}

	if args.modify || args.modifyAdd {
		return os.OpenFile(args.filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	return nil, errors.New("No correct options received to modify file")
}

// Map the actions passed in for logging
func actionMap(args Arguments) []string {
	var actionMap []string

	if args.create {
		actionMap = append(actionMap, "create")
	}
	if args.modifyReplace {
		actionMap = append(actionMap, "modifyReplace")
	}
	if args.modify {
		actionMap = append(actionMap, "modify")
	}
	if args.modifyAdd {
		actionMap = append(actionMap, "modifyAdd")
	}
	if args.delete {
		actionMap = append(actionMap, "delete")
	}
	return actionMap
}
