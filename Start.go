package serial2mqtt

import (
	"github.com/Jblew/serial2mqtt/mqttqueue"
	"github.com/Jblew/serial2mqtt/serialgateway"
)

// Start is starting all event loops
func Start(config Config) Serial2MQTT {
	queue := mqttqueue.Start(config.MQTTPublish)
	serialChan := make(chan serialgateway.Event)
	serialConfig := serialgateway.Config(config.Serial)
	gateway := serialgateway.Start(serialConfig, config.Clock, serialChan)

	serial2mqtt := Serial2MQTT{
		config,
		queue,
		gateway,
	}

	go serial2mqtt.runSerialEventHandlingLoop(serialChan)
	return serial2mqtt
}
