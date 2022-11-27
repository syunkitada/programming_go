package grpc_server

import (
	"io"
	"net"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"sample1/pkg/config"
	"sample1/pkg/grpc_pb"
)

type SimpleServer struct {
	Conf *config.Config
	Hub  *Hub
}

func NewSimpleServer(conf *config.Config) *SimpleServer {
	server := SimpleServer{
		Conf: conf,
		Hub:  newHub(),
	}
	go server.Hub.run()
	return &server
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

func (simpleServer *SimpleServer) Status(ctx context.Context, statusRequest *grpc_pb.StatusRequest) (*grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}

func (srv *SimpleServer) GetLogs(req *grpc_pb.GetLogRequest, stream grpc_pb.Simple_GetLogsServer) error {
	data := []string{
		"hoge", "piyo", "foo",
	}
	for _, d := range data {
		if err := stream.Send(&grpc_pb.GetLogStream{
			Msg: d,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (srv *SimpleServer) ReportLogs(stream grpc_pb.Simple_ReportLogsServer) error {
	for {
		reportLogStream, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&grpc_pb.ReportLogResponse{
				Msg: "END",
			})
		}
		if err != nil {
			return err
		}
		glog.Info(reportLogStream)
	}
}

func (srv *SimpleServer) Chat(stream grpc_pb.Simple_ChatServer) error {
	srv.Hub.register <- stream
	defer func() {
		srv.Hub.unregister <- stream
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		glog.Infof("Recieved: %v", in.Msg)

		srv.Hub.broadcast <- in

		// if err := stream.Send(in); err != nil {
		// 	return err
		// }
	}
}
