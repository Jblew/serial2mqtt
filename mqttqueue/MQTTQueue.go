package mqttqueue

import (
	"github.com/sheerun/queue"
)

// MQTTQueue is a thread safe, failproof mqtt sender
type MQTTQueue struct {
	mqttClient   MQTTClient
	messageQueue *queue.Queue
}
