package serialgateway

import (
	"io"
)

func (gateway *SerialGateway) nonblockingWrite(writer io.Writer) error {
	select {
	case payload := <-gateway.publishChan:
		return writeToSerial(writer, payload)
	default:
		return nil
	}
}

func writeToSerial(writer io.Writer, payload []byte) error {
	_, err := writer.Write(payload)
	if err != nil {
		return err
	}
	return nil
}
