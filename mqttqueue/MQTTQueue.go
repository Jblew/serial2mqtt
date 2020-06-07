package mqttqueue

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/sheerun/queue"
)

// MQTTQueue is a thread safe, failproof mqtt sender
type MQTTQueue struct {
	mqttClient   MQTT.Client
	messageQueue *queue.Queue
}
