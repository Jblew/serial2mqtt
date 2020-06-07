package serialgateway

func (gateway *SerialGateway) notifyError(err error) {
	evt := emptyEvent()
	evt.Error = &EventError{
		Time: gateway.clock.GetTime(),
		Err:  err,
	}
	gateway.outputChan <- evt
}

func (gateway *SerialGateway) notifyTextReceived(text string) {
	evt := emptyEvent()
	evt.TextReceived = &EventTextReceived{
		Time: gateway.clock.GetTime(),
		Text: text,
	}
	gateway.outputChan <- evt
}

func (gateway *SerialGateway) notifyFrameReceived(meta string, payload []byte) {
	evt := emptyEvent()
	evt.FrameReceived = &EventFrameReceived{
		Time:    gateway.clock.GetTime(),
		Meta:    meta,
		Payload: payload,
	}
	gateway.outputChan <- evt
}

func (gateway *SerialGateway) notifyConnected() {
	evt := emptyEvent()
	evt.Connected = &EventConnected{
		Time: gateway.clock.GetTime(),
	}
	gateway.outputChan <- evt
}

func (gateway *SerialGateway) notifyDisconnected(err error) {
	evt := emptyEvent()
	evt.Disconnected = &EventDisconnected{
		Time: gateway.clock.GetTime(),
		Err:  err,
	}
	gateway.outputChan <- evt
}

func (gateway *SerialGateway) notifyStale() {
	evt := emptyEvent()
	evt.Stale = &EventStale{
		Time: gateway.clock.GetTime(),
	}
	gateway.outputChan <- evt
}

func emptyEvent() Event {
	return Event{
		Error:         nil,
		TextReceived:  nil,
		FrameReceived: nil,
		Connected:     nil,
		Disconnected:  nil,
		Stale:         nil,
	}
}
