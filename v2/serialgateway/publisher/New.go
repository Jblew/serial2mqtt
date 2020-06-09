package publisher

import (
	"bufio"
	"io"
)

// New constructs Publisher out of a Writer
func New(baseWriter io.Writer) Publisher {
	writer := bufio.NewWriter(baseWriter)
	publisher := Publisher{
		writer,
	}
	return publisher
}
