package log

import (
	"fmt"

	"github.com/nats-io/nats.go"
	pb "github.com/samirgadkari/sidecar/protos/v1/messages"
)

type Logger struct {
	natsConn *nats.Conn
	topic    string
	header   *pb.Header
}

func NewLogger(natsConn *nats.Conn, header *pb.Header) *Logger {

	return &Logger{
		natsConn: natsConn,
		topic:    "search.v1.Log",
		header:   header,
	}
}

func (l *Logger) Log(s string, args ...interface{}) {

	str := fmt.Sprintf(s, args...)
	l.LogString(&str)
}

func (l *Logger) LogString(msg *string) error {

	// Print message to stdout
	fmt.Println(*msg)

	header := l.header
	header.MsgType = pb.MsgType_MSG_TYPE_LOG
	header.MsgId = 0

	l.natsConn.Publish(l.topic, []byte(*msg))

	return nil
}
