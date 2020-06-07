package serial2mqtt

import (
	"fmt"
	"time"
)

func publishToMQTT(message MQTTMessage) error {
	payload := *message.payload
	token := queue.mqttClient.Publish(message.topic, 1, false, payload)
	token.WaitTimeout(4 * time.Second)
	ok := token.Wait()
	if ok != true {
		return fmt.Errorf("MQTT timeout")
	}
	err := token.Error()
	if err != nil {
		return err
	}
	return nil
}
