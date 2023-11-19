package main

import (
	"log"
	"net"

	"github.com/msemjan/go-grpc/types"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	notificationServer := types.MyGRPCServer{}

	types.RegisterNotificationServiceServer(grpcServer, &notificationServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
