package main

import (
	"os"

	"github.com/bxcodec/go-clean-arch/forward"
	"github.com/bxcodec/go-clean-arch/internal/hook"
	mqttinfo "github.com/bxcodec/go-clean-arch/internal/repository/mqtt_info"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

var svcMqtt *mqtt.Server

func MqttInit() (res error) {

	svcMqtt = mqtt.New(nil)
	repo := mqttinfo.NewForwardRepository()
	svc := forward.NewService(repo)
	hook.NewForwardHandler(svcMqtt, svc)

	mqttHost := os.Getenv("DATABASE_HOST")
	tcpOpt := listeners.Config{
		Type:    "tcp",
		ID:      "tcp_mqtt",
		Address: mqttHost,
	}
	tcp := listeners.NewTCP(tcpOpt)
	svcMqtt.AddListener(tcp)

	return
}

func MqttServerStart() {
	svcMqtt.Serve()
}
