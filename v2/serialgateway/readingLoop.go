package serialgateway

import (
	"bufio"
	"strings"
	"time"
)

func (gateway *SerialGateway) readingLoop() error {
	bufReader := bufio.NewReader(gateway.currentConnection)
	for {
		err := gateway.doRead(bufReader)
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (gateway *SerialGateway) doRead(bufReader *bufio.Reader) error {
	if bufReader.Buffered() == 0 {
		return nil
	}
	return gateway.blockingRead(bufReader)
}

func (gateway *SerialGateway) blockingRead(bufReader *bufio.Reader) error {
	line, err := bufReader.ReadString('\n')
	if err != nil {
		return err
	} else if strings.HasPrefix(line, "FRAME") {
		err = gateway.readFrame(bufReader, line)
		if err != nil {
			return err
		}
	} else {
		gateway.notifyTextReceived(line)
	}
	return nil
}
