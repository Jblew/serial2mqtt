package serialgateway

import (
	"io"
	"log"
)

func (gateway *SerialGateway) nonblockingWrite(writer io.Writer) error {
	select {
	case payload := <-gateway.publishChan:
		writeToSerial(writer, payload)
		return nil
	default:
		return nil
	}
}

func writeToSerial(writer io.Writer, payload []byte) {
	_, err := writer.Write(payload)
	if err != nil {
		log.Printf("Error while writing to serial %v", err)
	}
}
