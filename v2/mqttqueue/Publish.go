package mqttqueue

// PublishBin adds binary message to publish queue
func (queue *MQTTQueue) PublishBin(topic string, payload []byte) {
	msg := MQTTMessage{
		topic:   topic,
		payload: &payload,
	}
	queue.messageQueue.Append(msg)
}

// PublishStr adds string message to publish queue
func (queue *MQTTQueue) PublishStr(topic string, payload string) {
	payloadBin := []byte(payload)
	msg := MQTTMessage{
		topic:   topic,
		payload: &payloadBin,
	}
	queue.messageQueue.Append(msg)
}
