package grpc_server

import (
	"net"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/go-samples/grpc-sample/pkg/config"
	"github.com/syunkitada/go-samples/grpc-sample/pkg/grpc_pb"
)

type SimpleServer struct {
	Conf *config.Config
}

func NewSimpleServer(conf *config.Config) *SimpleServer {
	server := SimpleServer{
		Conf: conf,
	}
	return &server
}

func (simpleServer *SimpleServer) Status(ctx context.Context, statusRequest *grpc_pb.StatusRequest) (*grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}

func (simpleServer *SimpleServer) Serv() error {
	grpcConfig := simpleServer.Conf.SimpleServer.Grpc

	lis, err := net.Listen("tcp", grpcConfig.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(
		simpleServer.Conf.Path(grpcConfig.CertFile),
		simpleServer.Conf.Path(grpcConfig.KeyFile),
	)
	if err != nil {
		return err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)

	grpc_pb.RegisterSimpleServer(grpcServer, simpleServer)
	glog.Infof("Serve %v", grpcConfig.Listen)
	grpcServer.Serve(lis)

	return nil
}
