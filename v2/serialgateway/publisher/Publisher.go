package publisher

import "bufio"

// Publisher sends data to serial
type Publisher struct {
	writer *bufio.Writer
}
