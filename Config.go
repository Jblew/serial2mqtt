package serial2mqtt

import (
	"serialsender/serial2mqtt/mqttqueue"
	"serialsender/serial2mqtt/serialgateway"
)

// Config for serial2mqtt gateway
type Config struct {
	Serial     serialgateway.SerialConfig
	Clock      serialgateway.Clock
	MQTTClient mqttqueue.MQTTClient
	Handlers   Handlers
	FrameTopic string
	LogTopic   string
}
