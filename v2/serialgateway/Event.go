package serialgateway

import (
	"fmt"
	"time"
)

// Event that occured on serial bus
type Event struct {
	Error         *EventError
	TextReceived  *EventTextReceived
	FrameReceived *EventFrameReceived
	Connected     *EventConnected
	Disconnected  *EventDisconnected
	Stale         *EventStale
}

// EventError - custom error happened
type EventError struct {
	Time time.Time
	Err  error
}

func (evt EventError) String() string {
	return fmt.Sprintf("EventError{%v}", evt.Err)
}

// EventTextReceived - text was received on serial bus
type EventTextReceived struct {
	Time time.Time
	Text string
}

func (evt EventTextReceived) String() string {
	return fmt.Sprintf("EventTextReceived{%v}", evt.Text)
}

// EventFrameReceived - frame was received on serial bus
type EventFrameReceived struct {
	Time    time.Time
	Payload []byte
	Meta    string
}

func (evt EventFrameReceived) String() string {
	return fmt.Sprintf("EventFrameReceived{len=%v; meta=%v}", len(evt.Payload), evt.Meta)
}

// EventConnected - serial bus has connected
type EventConnected struct {
	Time time.Time
}

func (evt EventConnected) String() string {
	return "EventConnected{}"
}

// EventDisconnected - serial bus has disconnected
type EventDisconnected struct {
	Time time.Time
	Err  error
}

func (evt EventDisconnected) String() string {
	return fmt.Sprintf("EventDisconnected{%v}", evt.Err)
}

// EventStale - serial bus is not sending data over bus
type EventStale struct {
	Time time.Time
}

func (evt EventStale) String() string {
	return "EventStale{}"
}

func (evt Event) String() string {
	evtStr := ""
	if evt.Error != nil {
		evtStr += evt.Error.String()
	}
	if evt.TextReceived != nil {
		evtStr += evt.TextReceived.String()
	}
	if evt.FrameReceived != nil {
		evtStr += evt.FrameReceived.String()
	}
	if evt.Connected != nil {
		evtStr += evt.Connected.String()
	}
	if evt.Disconnected != nil {
		evtStr += evt.Disconnected.String()
	}
	if evt.Stale != nil {
		evtStr += evt.Stale.String()
	}
	return evtStr
}
