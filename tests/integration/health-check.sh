#! /bin/bash
# Intended to be run from repository root

set -euo pipefail

# Import global variables + functions
source ./tests/integration/utils.sh

function make_request() {
    echo "Waiting 3 seconds"
    sleep 3

    echo "Send healthcheck request"
    curl http://localhost:8080/healthCheck

    echo "Waiting 2 seconds"
    sleep 2
}

main() {
    start_go_server
    start_rust_server
    make_request
    kill_servers
}

main
