package publisher

import (
	"bufio"
	"io"

	"github.com/sheerun/queue"
)

// New constructs Publisher out of a Writer
func New(baseWriter io.Writer) Publisher {
	writer := bufio.NewWriter(baseWriter)
	queue := queue.New()
	publisher := Publisher{
		writer,
		queue,
	}
	return publisher
}
