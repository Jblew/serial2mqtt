package serialgateway

import (
	"log"
	"time"
)

func (gateway *serialGateway) connectReadLoop() {
	for {
		err := gateway.openConnection()
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
