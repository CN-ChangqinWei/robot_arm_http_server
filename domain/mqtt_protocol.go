package domain

type MqttProtocol int

const (
	PROTO_HEALTH MqttProtocol = iota
	PROTO_ECHO
	PROTO_MOTOR
	NUM_OF_PROTO
)
