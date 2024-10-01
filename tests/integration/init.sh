#! /bin/bash
# Intended to be run from repository root
# Do 'make run' for both repos to initialize build caches

set -euo pipefail

# Import global variables + functions
source ./tests/integration/utils.sh

main() {
    start_go_server
    start_rust_server
}

main
