package serial2mqtt

import (
	"github.com/Jblew/serial2mqtt/v2/mqttqueue"
	"github.com/Jblew/serial2mqtt/v2/serialgateway"
)

// Serial2MQTT is main class for serial2mqtt lib
type Serial2MQTT struct {
	config  Config
	queue   mqttqueue.MQTTQueue
	gateway *serialgateway.SerialGateway
}
