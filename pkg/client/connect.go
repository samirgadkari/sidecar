package client

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/samirgadkari/sidecar/protos/v1/messages"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
)

type SC struct {
	client pb.SidecarClient
	header *pb.Header
	Logger *Logger
}

func InitSidecar(serviceName string, regParams *pb.RegistrationParams) *SC {

	sidecarServiceAddr := viper.GetString("sidecarServiceAddr")
	_, sidecar, err := Connect(serviceName, sidecarServiceAddr, regParams)
	if err != nil {
		fmt.Printf("Error connecting to client: %v\n", err)
		os.Exit(-1)
	}

	return sidecar
}

func Connect(serviceName string, serverAddr string, regParams *pb.RegistrationParams) (*grpc.ClientConn, *SC, error) {

	// var opts []grpc.DialOption

	fmt.Printf("serverAddr: %s\n", serverAddr)
	conn, err := grpc.Dial(serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Error creating GRPC channel\n:\terr: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("conn: %v\n", conn)

	client := pb.NewSidecarClient(conn)
	fmt.Printf("GRPC connection to sidecar created\n")

	header := &pb.Header{
		SrcServType: serviceName,
		DstServType: "sidecar",
		ServId:      []byte(""),
		MsgId:       0,
	}

	logger := NewLogger(&client, header)
	sc := &SC{client, header, logger}
	err = sc.Register(serviceName, regParams)
	if err != nil {
		fmt.Printf("Error registering client:\n\terr: %v\n", err)
		os.Exit(-1)
	}

	return conn, sc, nil
}

func (sc *SC) Register(serviceName string, regParams *pb.RegistrationParams) error {

	var rParams *pb.RegistrationParams

	rParams = regParams
	if rParams == nil {
		var circuitConsecutiveFailures uint32 = 3

		// Go Duration is in the time package: https://pkg.go.dev/time#Duration
		// Go Duration maps to protobuf Duration.
		// You can convert between them using durationpb:
		//   https://pkg.go.dev/google.golang.org/protobuf/types/known/durationpb
		debounceDelayDuration, err := time.ParseDuration("5s")
		if err != nil {
			sc.Logger.Log("Error creating Golang time duration.\nerr: %v\n", err)
			return err
		}
		debounceDelay := durationpb.New(debounceDelayDuration)

		var retryNum uint32 = 2

		retryDelayDuration, err := time.ParseDuration("2s")
		if err != nil {
			sc.Logger.Log("Error creating Golang time duration.\nerr: %v\n", err)
			return err
		}
		retryDelay := durationpb.New(retryDelayDuration)

		rParams = &pb.RegistrationParams{
			CircuitFailureThreshold: &circuitConsecutiveFailures,
			DebounceDelay:           debounceDelay,

			Retry: &pb.RetryBehavior{
				RetryNum:   &retryNum,
				RetryDelay: retryDelay,
			},
		}
	}

	header := sc.header
	header.MsgType = pb.MsgType_MSG_TYPE_REG
	header.MsgId = 0

	rMsg := &pb.RegistrationMsg{
		Header: header,

		ServiceName: serviceName,
		RegParams:   rParams,
	}

	rRsp, err := sc.client.Register(context.Background(), rMsg)
	sc.Logger.Log("Registration msg sent:\n\t%s\n", rMsg)
	if err != nil {
		sc.Logger.Log("Sending Registration caused error:\n\terr: %v\n", err)
		return err
	}

	sc.Logger.Log("Registration rsp received:\n\t%s\n", rRsp)

	sc.header.ServId = rRsp.Header.ServId

	return nil
}

func (sc *SC) Pub(topic string, data []byte) error {

	header := sc.header
	header.MsgType = pb.MsgType_MSG_TYPE_PUB
	header.MsgId = 0

	pubMsg := pb.PubMsg{
		Header: sc.header,
		Topic:  topic,
		Msg:    data,
	}

	pubRsp, err := sc.client.Pub(context.Background(), &pubMsg)
	sc.Logger.Log("Pub message sent: %s\n", pubMsg.String())
	if err != nil {
		sc.Logger.Log("Could not publish to topic: %s\n\tmessage:\n\tmsg: %v\n\terr: %v\n",
			topic, data, err)
		return err
	}
	sc.Logger.Log("Pub rsp received: %s\n", pubRsp)

	if pubRsp.RspHeader.Status != uint32(pb.Status_OK) {
		sc.Logger.Log("Error received while publishing to topic:\n\ttopic: %s\n\tmsg: %s\n\terr: %v\n",
			topic, data, err)
		return err
	}

	return nil
}

func (sc *SC) Sub(topic string, chanSize uint32) error {

	header := sc.header
	header.MsgType = pb.MsgType_MSG_TYPE_SUB
	header.MsgId = 0

	subMsg := pb.SubMsg{
		Header:   header,
		Topic:    topic,
		ChanSize: chanSize,
	}

	subRsp, err := sc.client.Sub(context.Background(), &subMsg)
	sc.Logger.Log("Sub message sent:\n\t%s\n", &subMsg)
	if err != nil {
		sc.Logger.Log("Could not subscribe to topic: %s\n\terr: %v\n",
			topic, err)
		return err
	}
	sc.Logger.Log("Sub rsp received:\n\t%s\n", subRsp)

	if subRsp.RspHeader.Status != uint32(pb.Status_OK) {
		sc.Logger.Log("Error received while publishing to topic:\n\ttopic: %s\n\terr: %v\n",
			topic, err)
		return err
	}

	return nil
}

func (sc *SC) Unsub(topic string) error {

	header := sc.header
	header.MsgType = pb.MsgType_MSG_TYPE_UNSUB
	header.MsgId = 0

	unsubMsg := pb.UnsubMsg{
		Header: header,
		Topic:  topic,
	}

	unsubRsp, err := sc.client.Unsub(context.Background(), &unsubMsg)
	sc.Logger.Log("Unsub message sent:\n\t%s\n", &unsubMsg)
	if err != nil {
		sc.Logger.Log("Could not unsubscribe from topic:\n\ttopic: %s\n\terr: %v\n", topic, err)
		return err
	}

	if unsubRsp.RspHeader.Status != uint32(pb.Status_OK) {
		sc.Logger.Log("Error received while unsubscribing to topic:\n\ttopic: %s\n\terr: %v", topic, err)
		return err
	}

	return nil
}

func (sc *SC) ProcessSubMsgs(topic string, chanSize uint32, f func(*pb.SubTopicResponse)) error {

	err := sc.Sub(topic, chanSize)
	if err != nil {
		return err
	}

	for {
		subTopicRsp, err := sc.Recv(topic)
		if err != nil {
			fmt.Printf("Error receiving from sidecar: %v\n", err)
			break
		}

		f(subTopicRsp)
	}

	return nil
}

func (sc *SC) Recv(topic string) (*pb.SubTopicResponse, error) {

	header := sc.header
	header.MsgId = 0

	recvMsg := pb.Receive{
		Header: header,
		Topic:  topic,
	}

	subTopicRsp, err := sc.client.Recv(context.Background(), &recvMsg)
	if err != nil {
		fmt.Printf("Could not receive from sidecar - err: %v\n", err)
		return nil, err
	}

	// Do not log message to NATS. This creates a loop.
	fmt.Printf("Client received from sidecar: %s\n", subTopicRsp.String())
	return subTopicRsp, nil
}
