package serial2mqtt

import (
	"github.com/Jblew/serial2mqtt/v2/mqttqueue"
	"github.com/Jblew/serial2mqtt/v2/serialgateway"
)

// Run is starting all event loops
func Run(config Config) {
	queue := mqttqueue.Start(config.MQTTPublish)
	serialChan := make(chan serialgateway.Event)
	serialConfig := serialgateway.Config(config.Serial)
	gateway := serialgateway.Start(serialConfig, config.Clock, serialChan)

	serial2mqtt := Serial2MQTT{
		config,
		queue,
		gateway,
	}

	serial2mqtt.runSerialEventHandlingLoop(serialChan)
}
