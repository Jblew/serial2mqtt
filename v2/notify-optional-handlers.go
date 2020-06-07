package serial2mqtt

import "log"

func (serial2MQTT *Serial2MQTT) notifyOptionalOnError(eventError error) {
	onError := serial2MQTT.config.Handlers.OnError
	if onError == nil {
		return
	}
	evt := ErrorEvent{
		Serial2MQTT: serial2MQTT,
		Err:         eventError,
	}
	err := onError(evt)
	if err != nil {
		log.Printf("Error in optional OnError handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnConnected() {
	onConnected := serial2MQTT.config.Handlers.OnConnected

	if onConnected == nil {
		return
	}
	err := onConnected(ConnectedEvent{
		Serial2MQTT: serial2MQTT,
	})
	if err != nil {
		log.Printf("Error in optional OnConnected handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnDisconnected(eventError error) {
	onDisconnected := serial2MQTT.config.Handlers.OnDisconnected
	if onDisconnected == nil {
		return
	}
	err := onDisconnected(DisconnectedEvent{
		Serial2MQTT: serial2MQTT,
		Err:         eventError,
	})
	if err != nil {
		log.Printf("Error in optional OnDisconnected handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnStale() {
	onStale := serial2MQTT.config.Handlers.OnStale
	if onStale == nil {
		return
	}
	err := onStale(StaleEvent{
		Serial2MQTT: serial2MQTT,
	})
	if err != nil {
		log.Printf("Error in optional OnStale handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnFrame(meta string, payload []byte) {
	onFrame := serial2MQTT.config.Handlers.OnFrame
	if onFrame == nil {
		return
	}
	err := onFrame(FrameEvent{
		Serial2MQTT: serial2MQTT,
		Meta:        meta,
		Payload:     payload,
	})
	if err != nil {
		log.Printf("Error in optional OnFrame handler %v", err)
	}
}

func (serial2MQTT *Serial2MQTT) notifyOptionalOnText(text string) {
	onText := serial2MQTT.config.Handlers.OnText
	if onText == nil {
		return
	}
	err := onText(TextEvent{
		Serial2MQTT: serial2MQTT,
		Text:        text,
	})
	if err != nil {
		log.Printf("Error in optional OnText handler %v", err)
	}
}
