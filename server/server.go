package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "zhaor.com/mapper-grpc/pkg/apis/dmi/v1"
)

type server struct{}

func (s *server) MapperRegister(ctx context.Context, in *pb.MapperRegisterRequest) (*pb.MapperRegisterReply, error) {
	fmt.Printf("mapperinfo: %+v\n", in.Mapper)
	return &pb.MapperRegisterReply{}, nil
}

func main() {
	lis, err := net.Listen("unix", "/tmp/a.sock")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDeviceManagerServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
