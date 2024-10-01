#! /bin/bash
# Use to kill process running on port 8080, when Go UserRequestHandler did not release the port

set -euo pipefail

PID=$(sudo lsof -i -P -n | grep LISTEN | grep 8080 | awk '{print $2}')
echo "Killing process with PID: $PID"
sudo kill -2 $PID