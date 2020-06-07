package mqttqueue

import (
	"github.com/sheerun/queue"
)

// MQTTQueue is a thread safe, failproof mqtt sender
type MQTTQueue struct {
	mqttPublish  func(topic string, payload []byte) error
	messageQueue *queue.Queue
}
