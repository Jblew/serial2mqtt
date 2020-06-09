package serial2mqtt

import (
	"fmt"
	"log"
	"time"

	"github.com/Jblew/serial2mqtt/v2/serialgateway"
)

func (serial2MQTT *Serial2MQTT) handleSerialEvent(event serialgateway.Event) error {
	handlers := []func(event serialgateway.Event) error{
		serial2MQTT.handleSerialError,
		serial2MQTT.handleSerialConnected,
		serial2MQTT.handleSerialDisconnected,
		serial2MQTT.handleSerialStale,
		serial2MQTT.handleSerialFrameReceived,
		serial2MQTT.handleSerialTextReceived,
	}

	for _, handler := range handlers {
		err := handler(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (serial2MQTT *Serial2MQTT) handleSerialError(event serialgateway.Event) error {
	if event.Error == nil {
		return nil
	}
	isoTime := event.Error.Time.Format(time.RFC3339)
	logMsg := fmt.Sprintf("[%s] Serial error: %v", isoTime, event.Error.Err)
	log.Print(logMsg)
	serial2MQTT.sendLog(logMsg)
	go serial2MQTT.notifyOptionalOnError(event.Error.Err)

	return nil
}

func (serial2MQTT *Serial2MQTT) handleSerialConnected(event serialgateway.Event) error {
	if event.Connected == nil {
		return nil
	}
	isoTime := event.Connected.Time.Format(time.RFC3339)
	logMsg := fmt.Sprintf("[%s] Serial connected", isoTime)
	log.Print(logMsg)
	serial2MQTT.sendLog(logMsg)
	go serial2MQTT.notifyOptionalOnConnected()

	return nil
}

func (serial2MQTT *Serial2MQTT) handleSerialDisconnected(event serialgateway.Event) error {
	if event.Disconnected == nil {
		return nil
	}

	isoTime := event.Disconnected.Time.Format(time.RFC3339)
	logMsg := fmt.Sprintf("[%s] Serial disconnected, err: %v", isoTime, event.Disconnected.Err)
	log.Print(logMsg)
	serial2MQTT.sendLog(logMsg)
	go serial2MQTT.notifyOptionalOnDisconnected(event.Disconnected.Err)

	return nil
}

func (serial2MQTT *Serial2MQTT) handleSerialStale(event serialgateway.Event) error {
	if event.Stale == nil {
		return nil
	}

	isoTime := event.Stale.Time.Format(time.RFC3339)
	logMsg := fmt.Sprintf("[%s] Serial is stale", isoTime)
	log.Print(logMsg)
	serial2MQTT.sendLog(logMsg)
	go serial2MQTT.notifyOptionalOnStale()

	return nil
}

func (serial2MQTT *Serial2MQTT) handleSerialFrameReceived(event serialgateway.Event) error {
	if event.FrameReceived == nil {
		return nil
	}

	isoTime := event.FrameReceived.Time.Format(time.RFC3339)
	subtopic := event.FrameReceived.Meta
	payload := event.FrameReceived.Payload

	logMsg := fmt.Sprintf("[%s] Serial frame received: len=%d, subtopic=%s\n", isoTime, len(payload), subtopic)
	log.Print(logMsg)
	serial2MQTT.sendLog(logMsg)
	serial2MQTT.sendFrame(subtopic, payload)
	go serial2MQTT.notifyOptionalOnFrame(event.FrameReceived.Meta, event.FrameReceived.Payload)

	return nil
}

func (serial2MQTT *Serial2MQTT) handleSerialTextReceived(event serialgateway.Event) error {
	if event.TextReceived == nil {
		return nil
	}

	isoTime := event.TextReceived.Time.Format(time.RFC3339)
	logMsg := fmt.Sprintf("[%s serial log] %s", isoTime, event.TextReceived.Text)
	log.Print(logMsg)
	serial2MQTT.sendLog(logMsg)
	go serial2MQTT.notifyOptionalOnText(event.TextReceived.Text)

	return nil
}
