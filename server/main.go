package main

import (
	"context"
	"google.golang.org/grpc"
    "net"
    "flag"
    "log"

	"fmt"

	pb "github.com/paul-kang-1/grpc-tests/proto"
)

type server struct {
	pb.UnimplementedArrayComparerServer
}

var port = flag.Int("port", 50054, "Server port number")

func (s *server) GetIntArray(ctx context.Context, in *pb.ArrayRequest) (*pb.IntArrayReply, error) {
	size := in.GetLength()
	if size <= 0 || size > 100 {
		return nil, fmt.Errorf("invalid size: %d. Size should be between 0 and 100 (Mb)", size)
	}
	array := make([]int32, int(250000*size))
	reply := &pb.IntArrayReply{F: array}
	return reply, nil
}

func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err) 
    }
    grpcServer := grpc.NewServer()
    pb.RegisterArrayComparerServer(grpcServer, &server{})
    err = grpcServer.Serve(lis)
    if err != nil {
        log.Fatalf("Failed to serve: %v", err) 
    }
}
