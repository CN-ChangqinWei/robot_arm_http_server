package main

import (
	"github.com/bxcodec/go-clean-arch/forward"
	"github.com/bxcodec/go-clean-arch/internal/hook"
	mqttinfo "github.com/bxcodec/go-clean-arch/internal/repository/mqtt_info"
	mqtt "github.com/mochi-mqtt/server/v2"
)

var svcMqtt *mqtt.Server

func MqttInit() (res error) {
	svcMqtt = mqtt.New(nil)
	repo := mqttinfo.NewForwardRepository()
	svc := forward.NewService(repo)
	hook.NewForwardHandler(svcMqtt, svc)

	return
}

func MqttServerStart() {

}
