package publisher

// AddToQueue adds to publish queue
func (publisher *Publisher) AddToQueue(payload []byte) {
	publisher.queue.Append(payload)
}
