curl -X POST http://localhost:9090/motor/control \
  -H "Content-Type: application/json" \
  -d '{
    "device_name": "dev0",
    "data": {
        "protocol": 2,
        "id":0,
        "powerOn": 1,
        "numAngel": 0,
        "denAngel": 180,
        "maxAngel": 360,
        "encode": 0,
        "spEncode": 0,
        "pwmNum": 0,
        "pwmDen": 0,
        "spNumAngel": 0,
        "spDenAngel": 0,
        "mode": 0
    }
}'
