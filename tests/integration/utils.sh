#! /bin/bash
# Intended to be run from repository root

set -euo pipefail

PID1=0
PID2=0
PROCESS_GROUP_ID=0

function start_go_server() {
    echo "Starting Golang server"
    make -C ./arbyhunter run > test1.log 2>&1 &
    PID1=$!
    PROCESS_GROUP_ID=$(ps -o pgid= -p $PID1)
}

function start_rust_server() {
    echo "Starting Rust server"
    make -C ./graph-engine run > test2.log 2>&1 &
    PID2=$!
}

function kill_servers() {
    # TODO - Not a clean exit, but I was finding that with `kill -2 $pid1`, the Golang app did not register a SIGINT signal as expected
    echo "Killing servers"
    kill -2 -$PROCESS_GROUP_ID
}
