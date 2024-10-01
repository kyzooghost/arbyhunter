#! /bin/bash
# Intended to be run from repository root

set -euo pipefail

# Import global variables + functions
source ./tests/integration/utils.sh

function make_request() {
    echo "Waiting 3 seconds"
    sleep 3

    echo "Send launchNodeAdaptor request"
    curl -X POST http://localhost:8080/launchNodeAdaptor \
        -d "{\"raw_url\": \"${ETH_RPC_URL}\", \"node_adaptor_type\": 0}" \
        -H "Content-Type: application/json"

    echo "Waiting 1 seconds"
    sleep 1

    echo "Send addPool request"
    curl -X POST http://localhost:8080/addPool \
        # -d "{\"raw_url\": \"${ETH_RPC_URL}\", \"node_adaptor_type\": 0}" \
        # -H "Content-Type: application/json"

    echo "Waiting 1 seconds"
    sleep 1
}

main() {
    start_go_server
    start_rust_server
    make_request
    kill_servers
}

main
