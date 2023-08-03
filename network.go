package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

func connectAndSendData(args Arguments) {
	// Let's get our of here if we don't have an endpoint to send to
	if args.endpoint == "" {
		return
	}

	// setup our connection
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	// make sure we close the connection
	defer conn.Close()

	// if we are sending a file
	if args.sendFile && args.filepath != "" {
		file, err := os.Open(args.filepath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer file.Close()
		bytesWritten, err := io.Copy(conn, file)
		if err != nil {
			fmt.Println("Error sending file:", err)
			return
		}
		logNetworkActivity(conn.LocalAddr().String(), conn.RemoteAddr().String(), bytesWritten, "tcp")
	} else {
		// If we included an endpoint, but no file, let's send the global profile info
		jsonData, err := json.MarshalIndent(globalProfile, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling global profile", err)
			return
		}
		bytesWritten, err := conn.Write(jsonData)
		if err != nil {
			fmt.Println("Error writing global profile to connection", err)
		}
		logNetworkActivity(conn.LocalAddr().String(), conn.RemoteAddr().String(), int64(bytesWritten), "tcp")
	}
}
