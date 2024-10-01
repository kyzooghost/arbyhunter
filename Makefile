# Run make run for servers for first time on a local machine
init :; . ./tests/integration/init.sh
test-hc :; . ./tests/integration/health-check.sh
test-pool :; . ./tests/integration/add-pool.sh
fix-port :; . ./scripts/fix-port.sh