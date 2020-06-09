package serialgateway

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func (gateway *SerialGateway) readingLoop() error {
	bufReader := bufio.NewReader(*gateway.currentConnection)
	for {
		err := gateway.doRead(bufReader)
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (gateway *SerialGateway) doRead(bufReader *bufio.Reader) error {
	line, err := bufReader.ReadString('\n')
	if err != nil {
		return err
	} else if strings.HasPrefix(line, "FRAME") {
		err = gateway.readFrame(bufReader, line)
		if err != nil {
			return err
		}
	} else {
		gateway.notifyTextReceived(line)
	}
	time.Sleep(10 * time.Millisecond)
	return nil
}

func (gateway *SerialGateway) readFrame(bufReader *bufio.Reader, rawMeta string) error {
	head, err := parseFrameHead(rawMeta)
	if err != nil {
		gateway.notifyError(err)
		return nil
	}

	payload := make([]byte, head.payloadLength)
	_, err = io.ReadFull(bufReader, payload)
	if err != nil {
		return err
	}
	gateway.notifyFrameReceived(head.meta, payload)
	return nil
}

func parseFrameHead(rawHead string) (frameHead, error) {
	rawHeadNoNewline := rawHead[:len(rawHead)-1]
	metaParts := strings.Split(rawHeadNoNewline, ":")
	if len(metaParts) < 3 {
		return frameHead{}, fmt.Errorf("Invalid frame head %v", rawHead)
	}

	payloadLength, err := strconv.ParseUint(metaParts[1], 10, 16)
	if err != nil {
		return frameHead{}, fmt.Errorf("Invalid frame length %v", err)
	}

	meta := metaParts[2]
	return frameHead{
		payloadLength: uint(payloadLength),
		meta:          meta,
	}, nil
}

type frameHead struct {
	payloadLength uint
	meta          string
}
