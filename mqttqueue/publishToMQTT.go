package mqttqueue

func (queue *MQTTQueue) publishToMQTT(message MQTTMessage) error {
	payload := *message.payload
	err := queue.mqttClient.Publish(message.topic, payload)
	if err != nil {
		return err
	}
	return nil
}
