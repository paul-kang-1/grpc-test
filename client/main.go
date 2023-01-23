package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	pb "github.com/paul-kang-1/grpc-test/proto"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = flag.Int("port", 50054, "Server port number")
var size = flag.Int("size", 10000, "Array size")
var target = flag.Arg(0)

func main () {
    flag.Parse()
    conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        logger.Fatalf("Failed to dial port %d, err: %v", *port, err)
    }    
    client := pb.NewArrayComparerClient(conn)
    req := &pb.ArrayRequest{Length: int32(*size)}
    start := time.Now()
    switch target {
    case "int":
        _, err = client.GetIntArray(context.Background(), req) 
    case "user":
        fallthrough
    default:
        _, err = client.GetUserArray(context.Background(), req)
    }
    if err != nil {
        logger.Fatalf("Request Failed: %v", err)
    }    
    logger.Infof("Got response of %d elements in %d ms", *size, time.Since(start).Milliseconds())
    defer conn.Close()
}
