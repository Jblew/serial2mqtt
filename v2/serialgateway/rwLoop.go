package serialgateway

import (
	"bufio"
	"strings"
	"time"
)

func (gateway *SerialGateway) rwLoop() error {
	bufReader := bufio.NewReader(gateway.currentConnection)
	// TODO make separate loop for writing (so that writing does not wait for reading)
	for {
		gateway.nonblockingWrite(gateway.currentConnection)
		err := gateway.blockingRead(bufReader)
		if err != nil {
			return err
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
