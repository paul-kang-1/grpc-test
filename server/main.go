package main

import (
	"context"
	"flag"
	"net"
	"time"

	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"fmt"

	pb "github.com/paul-kang-1/grpc-test/proto"
)

type server struct {
	pb.UnimplementedArrayComparerServer
}

var port = flag.Int("port", 50054, "Server port number")

func (s *server) GetIntArray(ctx context.Context, in *pb.ArrayRequest) (*pb.IntArrayReply, error) {
	size := in.GetLength()
	if size <= 0 {
		return nil, fmt.Errorf("invalid size: %d. Size should be larger than 0", size)
	}
	array := make([]int32, int(size))
    for idx := range array {
        array[idx] = 123
    }
	reply := &pb.IntArrayReply{F: array}
	return reply, nil
}

func (s *server) GetUserArray(ctx context.Context, in *pb.ArrayRequest) (*pb.UserArrayReply, error) {
    size := int(in.GetLength())
	if size <= 0 {
		return nil, fmt.Errorf("invalid size: %d. Size should be larger than 0", size)
	}
    array := make([]*pb.UserResponse, size)
    start := time.Now()
    // Generate user data
    for i := 0; i < size; i++ {
        name := fmt.Sprintf("TestUser%d", i)
        email := fmt.Sprintf("%s@gmail.com", name)
        user := &pb.UserResponse{
            Id: int64(i), 
            Username: name,
            DisplayName: name,
            Profile: &name,
            Email: &email,
        }
        array[i] = user
    }
    logger.Infof("Packed userlist in %d ms", time.Since(start).Milliseconds())
    reply := &pb.UserArrayReply{Users: array}
    return reply, nil
} 

func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
    if err != nil {
        logger.Fatalf("Failed to listen: %v", err) 
    }
    grpcServer := grpc.NewServer()
    pb.RegisterArrayComparerServer(grpcServer, &server{})
    logger.Infof("Running server on %d", *port)
    err = grpcServer.Serve(lis)
    if err != nil {
        logger.Fatalf("Failed to serve: %v", err) 
    }
}
