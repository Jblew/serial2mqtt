package serial2mqtt

import "time"

// Clock allows to obtain exact time of received events
type Clock interface {
	GetTime() time.Time
}
