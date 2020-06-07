package mqttqueue

import (
	"github.com/sheerun/queue"
)

// MQTTQueue is a thread safe, failproof mqtt sender
type MQTTQueue struct {
	mqttPublish  MQTTPublish
	messageQueue *queue.Queue
}
