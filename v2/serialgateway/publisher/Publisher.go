package publisher

import (
	"bufio"

	"github.com/sheerun/queue"
)

// Publisher sends data to serial
type Publisher struct {
	writer *bufio.Writer
	queue  *queue.Queue
}
