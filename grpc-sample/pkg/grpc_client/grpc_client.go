package grpc_client

import (
	"errors"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/go-samples/grpc-sample/pkg/config"
	"github.com/syunkitada/go-samples/grpc-sample/pkg/grpc_pb"
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
