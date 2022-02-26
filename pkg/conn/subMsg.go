package conn

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/samirgadkari/sidecar/protos/v1/messages"
)

const (
	msgChSize = 10
)

type Subs struct {
	natsMsgs chan *nats.Msg
	done     chan struct{}
	msgId    uint32
	natsConn *Conn
	header   *messages.Header
}

func InitSubs(natsConn *Conn, srv *Server) {

	natsMsgs := make(chan *nats.Msg, msgChSize)
	done := make(chan struct{})

	srv.Subs = &Subs{

		natsMsgs: natsMsgs,
		done:     done,
		msgId:    1,
		natsConn: natsConn,
	}
}

func (subs *Subs) Subscribe(in *messages.SubMsg) (*messages.SubMsgResponse, error) {

	fmt.Printf("Received SubMsg: %v\n", in)
	topic := in.GetTopic()

	subs.natsConn.Subscribe(topic, func(m *nats.Msg) {
		fmt.Printf("Received msg from NATS: %v\n", m)
		subs.natsMsgs <- m
	})

	srcHeader := in.GetHeader()
	header := messages.Header{
		DstServType: srcHeader.GetSrcServType(),
		SrcServType: srcHeader.GetDstServType(),
		ServId:      srcHeader.GetServId(),
		MsgId:       0,
	}
	subs.header = &header

	rspHeader := messages.ResponseHeader{

		Status: uint32(messages.Status_OK),
	}

	subMsgRsp := messages.SubMsgResponse{

		Header:    &header,
		RspHeader: &rspHeader,
		Msg:       "OK",
	}

	return &subMsgRsp, nil
}

func RecvFromNATS(srv *Server) (*messages.SubTopicResponse, error) {

	m := <-srv.Subs.natsMsgs
	fmt.Printf("Got msg from NATS server:\n\t%#v\n", m)

	return &messages.SubTopicResponse{
		Header: srv.Subs.header,
		Topic:  m.Subject,
		Msg:    m.Data,
	}, nil
}
