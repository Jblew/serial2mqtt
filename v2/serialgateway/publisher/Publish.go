package publisher

// Publish sends data to serial and flushes
func (publisher *Publisher) Publish(payload []byte) error {

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
