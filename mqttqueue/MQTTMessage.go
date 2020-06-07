package mqttqueue

import "fmt"

// MQTTMessage is a message pushed to publishing queue
type MQTTMessage struct {
	topic   string
	payload *[]byte
}

func (msg MQTTMessage) String() string {
	return fmt.Sprintf("MQTTMessage{to=%v}", msg.topic)
}
