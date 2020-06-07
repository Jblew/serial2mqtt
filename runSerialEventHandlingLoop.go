package serial2mqtt

import (
	"log"
	"serialsender/serial2mqtt/serialgateway"
	"time"
)

func (serial2MQTT *Serial2MQTT) runSerialEventHandlingLoop(serialChan chan serialgateway.Event) error {
	for {
		time.Sleep(2 * time.Millisecond)
		evt := <-serialChan
		serial2MQTT.handleSerialEvent(evt)
	}
}

func (serial2MQTT *Serial2MQTT) handleEvent(event serialgateway.Event) {
	err := serial2MQTT.handleSerialEvent(event)
	if err != nil {
		log.Printf("Error while handling serial event %v", err)
		go serial2MQTT.notifyOptionalOnError(err)
	}
}
