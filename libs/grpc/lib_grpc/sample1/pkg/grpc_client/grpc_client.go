package grpc_client

import (
	"bufio"
	"errors"
	"io"
	"os"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"sample1/pkg/config"
	"sample1/pkg/grpc_pb"
)

type SimpleClient struct {
	Conf               *config.Config
	CaFilePath         string
	ServerHostOverride string
	Targets            []string
	Timeout            time.Duration
}

func NewSimpleClient(conf *config.Config) *SimpleClient {
	simpleClient := SimpleClient{
		Conf:               conf,
		CaFilePath:         conf.Path(conf.SimpleServer.Grpc.CaFile),
		ServerHostOverride: conf.SimpleServer.Grpc.ServerHostOverride,
		Targets:            conf.SimpleServer.Grpc.Targets,
		Timeout:            time.Duration(conf.SimpleServer.Grpc.Timeout) * time.Second,
	}
	return &simpleClient
}

func (client *SimpleClient) NewClientConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	for _, target := range client.Targets {
		creds, credsErr := credentials.NewClientTLSFromFile(client.CaFilePath, client.ServerHostOverride)
		if credsErr != nil {
			glog.Warning("Failed to create TLS credentials %v", credsErr)
			continue
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

		conn, err := grpc.Dial(target, opts...)
		if err != nil {
			glog.Warning("fail to dial: %v", err)
			continue
		}

		return conn, nil
	}

	return nil, errors.New("Failed NewGrpcConnection")
}

func (client *SimpleClient) Status() (*grpc_pb.StatusReply, error) {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	rpcClient := grpc_pb.NewSimpleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), client.Timeout)
	defer cancel()

	statusResponse, err := rpcClient.Status(ctx, &grpc_pb.StatusRequest{})
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", client, err)
		return nil, err
	}

	return statusResponse, nil
}

func (client *SimpleClient) GetLogs() {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return
	}

	rpcClient := grpc_pb.NewSimpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &grpc_pb.GetLogRequest{
		Msg: "hoge",
	}
	stream, err := rpcClient.GetLogs(ctx, req)
	if err != nil {
		glog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			glog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		glog.Info(feature)
	}
}

func (client *SimpleClient) ReportLogs() {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return
	}

	rpcClient := grpc_pb.NewSimpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := rpcClient.ReportLogs(ctx)
	if err != nil {
		glog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}

	msgs := []string{
		"hoge", "piyo",
	}
	for _, msg := range msgs {
		if err = stream.Send(&grpc_pb.ReportLogStream{
			Msg: msg,
		}); err != nil {
			glog.Error(err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		glog.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	glog.Infof("Route summary: %v", reply)
}

func (client *SimpleClient) Chat() {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return
	}

	rpcClient := grpc_pb.NewSimpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	stream, err := rpcClient.Chat(ctx)
	if err != nil {
		glog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				glog.Fatalf("Failed to receive a note : %v", err)
			}
			glog.Infof("Got message %v", in)
		}
	}()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if err := stream.Send(&grpc_pb.ChatStream{
			Msg: s.Text(),
		}); err != nil {
			glog.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
