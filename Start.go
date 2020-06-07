package serial2mqtt

import (
	"serialsender/serial2mqtt/mqttqueue"
	"serialsender/serial2mqtt/serialgateway"
)

// Start is starting all event loops
func Start(config Config) Serial2MQTT {
	queue := mqttqueue.Start(mqttClient)
	serialChan := make(chan serialgateway.Event)
	gateway := serialgateway.Start(config.Serial, app.clock, serialChan)

	serial2mqtt := Serial2MQTT{
		config,
		queue,
		gateway,
	}

	go serial2mqtt.runSerialEventHandlingLoop(serialChan)
	return serial2mqtt
}
