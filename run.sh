#!/bin/bash



# Find and kill processes using port 8082
echo "Searching for processes using port 8082..."
PID_LIST=$(lsof -t -i :8082)

if [ -z "$PID_LIST" ]; then
    echo "No processes found using port 8082."
else
    echo "Killing processes with PIDs: $PID_LIST"
    kill -9 $PID_LIST
    echo "Processes killed."
fi

go run main.go
exit 0
