package mqttqueue

import (
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/sheerun/queue"
)

// Start instationates and starts MQTTQueue
func Start(mqttClient MQTT.Client) MQTTQueue {
	messageQueue := queue.New()
	queue := MQTTQueue{
		mqttClient,
		messageQueue,
	}

	go queue.loopOverQueue()

	return queue
}

func (queue *MQTTQueue) loopOverQueue() {
	for {
		time.Sleep(5 * time.Millisecond)
		msg := queue.messageQueue.Pop().(MQTTMessage)
		go queue.publishOrPushBack(msg)
	}
}
