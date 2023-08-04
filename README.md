### Red Canary Take Home -- Andrew Kelley

The goal of this project is to generate telemetry for an Endpoint Detection and Response agent to process.

We need to be able to do the following:
 - Process Creation
 - File Creation, Modification, and Deletion
 - Establish a network connection and transmit data

I chose to use Golang for this assignment for its ease of portability between systems. Also, I've just been wanting to experiment with it!

The executable is called ```red-canary-endpoint-tester```. If you pass the ```-help``` flag, you will get a list of options that you are able to pass.

Each time you run the executable, you are able to do each of the following with the passed in options: call another executable, create/modify/delete a file, and send something over the network. 

The files are separated out by use case. A rough path through the application is as follows:
```
main.go
log.go
profile.go
flags.go
handle_process.go
handle_files.go
network.go
```

Overview of Flags:
```
String Flags:
 -ep - the endpoint to connect to
 -ea - arguments to pass to the executable defined in -ef, form is -ea='-arg1|-arg2=?', with arguments separated by a pipe
 -ef - the path to the executable file to run
 -fp - the file path to create/update/delete
 -fs - The string to add to the file
 -lf - the path to the log file, defaults to ./logFile.json

Boolean Flags:
  -c - Create a new file at filepath -fp
  -d - Delete the file at filepath -fp
  -m - Modify the file at filepath -fp
  -ma - Modify the file at filepath -fp by adding to the end of the file
  -mr - Modify the file at filepath -fp by truncating and adding to the file
  -sf - If there is a file at -fp, and this flag is true, send the file to the endpoint specified in -ep
```

A couple notes:
1. The logs are in JSON format. I thought this would be the easiest format to consume
2. If we are not sending a file to the endpoint specified in -ep, we will send profile information in a json format, this is the information in the ```Profile``` struct

A few sample commands (make sure you have a server accepting requests at whatever -ep you specify):
```
./red-canary-endpoing-tester

./red-canary-endpoint-tester -ep=localhost:8080 -c -fp="tester.txt" -fs="This is a test"

./red-canary-endpoint-tester -m -ma -fp="tester.txt" -fs="Another string to add"

./red-canary-endpoint-tester -ep=localhost:8080 -d -fp="tester.txt" -fs="test string" 

```