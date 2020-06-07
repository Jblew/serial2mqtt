package mqttqueue

import (
	"log"
	"time"
)

func (queue *MQTTQueue) publishOrPushBack(msg MQTTMessage) {
	payload := *msg.payload
	err := queue.mqttPublish(msg.topic, payload)
	if err != nil {
		log.Printf("Error while sending message(%v): %v. Pushing back after 500ms", msg, err)
		time.Sleep(500 * time.Millisecond)
		queue.messageQueue.Append(msg)
	}
}
