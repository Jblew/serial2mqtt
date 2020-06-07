package serial2mqtt

// Handlers for optional logic
type Handlers struct {
	OnError        func(ErrorEvent)
	OnConnected    func(ConnectedEvent)
	OnDisconnected func(DisconnectedEvent)
	OnStale        func(StaleEvent)
	OnFrame        func(FrameEvent)
	OnText         func(TextEvent)
}

// ErrorEvent is optional handler for errors
type ErrorEvent struct {
	Err         error
	Serial2MQTT Serial2MQTT
}

// ConnectedEvent is optional handler for connection
type ConnectedEvent struct {
	Serial2MQTT Serial2MQTT
}

// DisconnectedEvent is optional handler for disconnection
type DisconnectedEvent struct {
	Err         error
	Serial2MQTT Serial2MQTT
}

// StaleEvent is optional handler for stale condition
// Stale condition occurs when serial stops sending data
type StaleEvent struct {
	Serial2MQTT Serial2MQTT
}

// FrameEvent is optional handler fired alongside publishing frame
type FrameEvent struct {
	Meta        string
	payload     []byte
	Serial2MQTT Serial2MQTT
}

// TextEvent is optional handler fired alongside publishing serial text
type TextEvent struct {
	Text        string
	Serial2MQTT Serial2MQTT
}
