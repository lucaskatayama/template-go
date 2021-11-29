package grpc

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lucaskatayama/hexgo/internal/ui/grpc/pb"
	"github.com/lucaskatayama/hexgo/pkg/log"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func Run() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTodoServer(s, &pb.TodoServerImpl{})
	log.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
