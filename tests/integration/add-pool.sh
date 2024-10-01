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

    # Use WBTC/ETH pool on Uniswap on Ethereum network
    # See https://eips.ethereum.org/EIPS/eip-7528 for ETH native address
    echo "Send addPool request"
    curl -X POST http://localhost:8080/addPool \
        -H "Content-Type: application/json" \
        --data-binary @- << EOF
        {
            "node_adaptor_type": 0,
            "protocol_adaptor_type": 0,
            "pool_address": "0xCBCdF9626bC03E24f779434178A73a0B4bad62eD",
            "assets": [
                {
                "ticker": "ETH",
                "address": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
                },
                {
                "ticker": "WBTC",
                "address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"
                }
            ]
        }
EOF
        
    echo "Waiting 5 seconds"
    sleep 5
}

main() {
    start_go_server
    start_rust_server
    make_request
    kill_servers
}

main
