curl -X POST http://localhost:9090/robot/control \
  -H "Content-Type: application/json" \
  -d '{
    "device_name": "dev0",
    "data": {
        "protocol": 3,
        "x": 0,
        "y": 100,
        "z": -100
    }
}'