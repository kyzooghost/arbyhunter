Target 1 - addPool flow to add Uniswap pairs using Ethereum node HTTP endpoint

Test query
```bash
curl -X POST http://localhost:8080/launchNodeAdaptor \
     -d '{"raw_url": "hello", "node_adaptor_type": 1}' \
     -H "Content-Type: application/json"
```