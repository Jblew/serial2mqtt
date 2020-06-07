package app

import (
	"fmt"
)

func (serial2MQTT *Serial2MQTT) sendFrame(subtopic string, payload []byte) {
	topic := fmt.Sprintf("%s/%s", serial2MQTT.config.FrameTopic, subtopic)
	serial2MQTT.queue.PublishBin(topic, payload)
}

func (serial2MQTT *Serial2MQTT) sendLog(msg string) {
	serial2MQTT.queue.PublishStr(serial2MQTT.config.LogTopic, msg)
}
