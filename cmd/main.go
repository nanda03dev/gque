package main

import (
	"log"
	"net"

	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/config"
	"github.com/nanda03dev/gque/grpc_handler"
	pb "github.com/nanda03dev/gque/proto"
	"github.com/nanda03dev/gque/services"
	"google.golang.org/grpc"
)

const (
	// Port for gRPC server to listen to
	GQUE_PORT = ":5456"
)

func main() {
	config.LoadConfig()

	common.InitializeChannels()

	config.SetupDatabase()

	AppServices := services.InitializeServices()

	go func() {
		lis, err := net.Listen("tcp", GQUE_PORT)

		if err != nil {
			log.Fatalf("failed connection: %v", err)
		}

		s := grpc.NewServer()

		pb.RegisterGqueServiceServer(s, &grpc_handler.GqueServer{Services: AppServices})

		log.Printf("server listening at %v", lis.Addr())

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server: %v", err)
		}
	}()

	// Ensure the main goroutine doesn't exit immediately
	select {}
}
