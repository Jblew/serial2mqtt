package serial2mqtt

import (
	"time"
)

// Config for serial2mqtt gateway
type Config struct {
	Serial      SerialConfig
	Clock       Clock
	MQTTPublish MQTTPublish
	Handlers    Handlers
	FrameTopic  string
	LogTopic    string
}

// SerialConfig - config of serial port
type SerialConfig struct {
	PortNames       string
	BaudRate        uint
	DataBits        uint
	StopBits        uint
	MinimumReadSize uint
}

// MQTTPublish is an interface that allows publishing MQTT messages
type MQTTPublish func(topic string, payload []byte) error

// Clock allows to obtain exact time of received events
type Clock interface {
	GetTime() time.Time
}
