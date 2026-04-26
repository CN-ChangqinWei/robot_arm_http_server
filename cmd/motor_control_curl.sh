curl -X POST http://localhost:9090/motor/control \
  -H "Content-Type: application/json" \
  -d '{
    "device_name": "dev0",
    "data": {
        "protocol": 2,
        "powerOn": 1,
        "numAngel": 100,
        "denAngel": 180,
        "maxAngel": 360,
        "encode": 0,
        "spEncode": 0,
        "pwmNum": 50,
        "pwmDen": 100,
        "spNumAngel": 0,
        "spDenAngel": 0,
        "mode": 0
    }
}'
