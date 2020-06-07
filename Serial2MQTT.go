package serial2mqtt

import (
	"github.com/Jblew/serial2mqtt/mqttqueue"
	"github.com/Jblew/serial2mqtt/serialgateway"
)

// Serial2MQTT is main class for serial2mqtt lib
type Serial2MQTT struct {
	config  Config
	queue   mqttqueue.MQTTQueue
	gateway serialgateway.SerialGateway
}
