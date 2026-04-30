curl -X POST http://localhost:8080/robot/control \
  -H "Content-Type: application/json" \
  -d '{
    "device_name": "dev0",
    "data": {
        "protocol": 3,
        "x": 100.5,
        "y": 200.0,
        "z": 50.0
    }
}'