package serialgateway

import (
	"bufio"
	"strings"
	"time"
)

func (gateway *SerialGateway) rwLoop() error {
	readErrorChan := make(chan error)
	go gateway.readLoop(readErrorChan)
	for {
		err := gateway.nonblockingWrite(gateway.currentConnection)
		if err != nil {
			return err
		}
		select {
		case readError := <-readErrorChan:
			return readError
		default:
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (gateway *SerialGateway) readLoop(readErrorChan chan error) {
	bufReader := bufio.NewReader(gateway.currentConnection)
	for {
		gateway.nonblockingWrite(gateway.currentConnection)
		err := gateway.blockingRead(bufReader)
		if err != nil {
			readErrorChan <- err
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
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
