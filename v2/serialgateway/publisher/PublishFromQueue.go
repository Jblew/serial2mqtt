package publisher

import "log"

// PublishFromQueue sends data to serial and flushes
func (publisher *Publisher) PublishFromQueue() {
	if publisher.queue.Length() == 0 {
		return
	}
	item := publisher.queue.Pop()
	payload := item.(payload)

	err := publisher.doPublish(payload.data)
	if err != nil {
		log.Printf("Error while publishing from queue to serial: %v", err)
	}
	publisher.PublishFromQueue()
}

func (publisher *Publisher) doPublish(payload []byte) error {
	_, err := publisher.writer.Write(payload)
	if err != nil {
		return err
	}

	err = publisher.writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
