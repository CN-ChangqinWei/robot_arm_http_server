package main

import (
	mqtt "github.com/mochi-mqtt/server/v2"
)

var server *mqtt.Server

func MqttInit() (res error) {
	server = mqtt.New(nil)

	return
}
