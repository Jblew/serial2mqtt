package serialgateway

import (
	"bufio"
	"io"
	"log"
	"strings"
	"time"
)

func (gateway *SerialGateway) rwLoop() error {
	bufReader := bufio.NewReader(gateway.currentConnection)
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

func (gateway *SerialGateway) nonblockingWrite(writer io.Writer) error {
	select {
	case payload := <-gateway.publishChan:
		log.Printf("ATTEMPT TO WRITE len=%d", len(payload))
		return nil
	default:
		return nil
	}
}
