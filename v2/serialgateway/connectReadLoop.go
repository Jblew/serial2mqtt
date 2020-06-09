package serialgateway

import (
	"log"
	"time"
)

func (gateway *SerialGateway) connectReadLoop() {
	for {
		err := gateway.openConnection()
		log.Printf("<><><> Current connection in connectReadLoop %v", *gateway.currentConnection)
		if err != nil {
			gateway.notifyError(err)
			log.Printf("Error connecting serial %v. Waiting 10 seconds", err)
			time.Sleep(10 * time.Second)
			continue
		}
		gateway.notifyConnected()

		err = gateway.readingLoop()
		if err != nil {
			gateway.notifyDisconnected(err)
		}
	}
}
