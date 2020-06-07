package serialgateway

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/jacobsa/go-serial/serial"
)

func (gateway *serialGateway) openConnection() error {
	config := gateway.config
	portNames := getPortNames(config.PortNames)

	currentConnection, err := tryAllPorts(config, portNames)
	if err != nil {
		return err
	}

	gateway.currentConnection = currentConnection

	return nil
}

func getPortNames(portNames string) []string {
	return strings.Split(portNames, ",")
}

func tryAllPorts(config Config, portNames []string) (io.ReadWriteCloser, error) {
	var lastError error = fmt.Errorf("Could not connect to any port. No errors reported. Did you specify any ports?")

	for _, portName := range portNames {
		log.Printf("Trying to connect to port %s", portName)
		connection, err := tryPort(config, portName)
		if err != nil {
			log.Printf("Connection to port %s failed: %v", portName, err)
			lastError = err
		} else {
			return connection, nil
		}
	}

	return nil, lastError
}

func tryPort(config Config, portName string) (io.ReadWriteCloser, error) {
	openOptions := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        config.BaudRate,
		DataBits:        config.DataBits,
		StopBits:        config.StopBits,
		MinimumReadSize: config.MinimumReadSize,
	}
	conn, err := serial.Open(openOptions)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
