package serial2mqtt

import "log"

func (serial2MQTT *Serial2MQTT) notifyOptionalOnError(eventError error) {
	if serial2MQTT.config.handlers.OnError == nil {
		return
	}
	err := serial2MQTT.config.handlers.OnError(ErrorEvent{
		Serial2MQTT: serial2MQTT,
		Err:         eventError,
	})
	if err != nil {
		log.Printf("Error in optional OnError handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnConnected() {
	if serial2MQTT.config.handlers.OnConnected == nil {
		return
	}
	err := serial2MQTT.config.handlers.OnConnected(ConnectedEvent{
		Serial2MQTT: serial2MQTT,
	})
	if err != nil {
		log.Printf("Error in optional OnConnected handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnDisconnected(eventError error) {
	if serial2MQTT.config.handlers.OnDisconnected == nil {
		return
	}
	err := serial2MQTT.config.handlers.OnDisconnected(DisconnectedEvent{
		Serial2MQTT: serial2MQTT,
		Err:         eventError,
	})
	if err != nil {
		log.Printf("Error in optional OnDisconnected handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnStale() {
	if serial2MQTT.config.handlers.OnStale == nil {
		return
	}
	err := serial2MQTT.config.handlers.OnStale(StaleEvent{
		Serial2MQTT: serial2MQTT,
	})
	if err != nil {
		log.Printf("Error in optional OnStale handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnFrame(Meta string, payload []byte) {
	if serial2mqtt.config.handlers.OnFrame == nil {
		return
	}
	err := serial2mqtt.config.handlers.OnFrame(FrameEvent{
		Serial2MQTT: serial2MQTT,
		Meta:        meta,
		Payload:     payload,
	})
	if err != nil {
		log.Printf("Error in optional OnFrame handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnText(text string) {
	if serial2mqtt.config.handlers.OnText == nil {
		return
	}
	err := serial2mqtt.config.handlers.OnText(TextEvent{
		Serial2MQTT: serial2MQTT,
		Text:        text,
	})
	if err != nil {
		log.Printf("Error in optional OnText handler %v", err)
	}
}
