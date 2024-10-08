package main

import (
	"log"
	"net"
	"os"

	"github.com/nanda03dev/gque/common"
	"github.com/nanda03dev/gque/config"
	"github.com/nanda03dev/gque/grpc_handler"
	pb "github.com/nanda03dev/gque/proto"
	"github.com/nanda03dev/gque/services"
	"github.com/nanda03dev/gque/workers"
	"google.golang.org/grpc"
)

var (
	GQUE_PORT = "5456"
)

func main() {
	if port := os.Getenv("GQUE_PORT"); port != "" {
		GQUE_PORT = port
	}

	config.LoadConfig()

	common.InitializeChannels()

	config.SetupDatabase()

	AppServices := services.InitializeServices()

	go workers.InitiateWorker()

	go func() {
		lis, err := net.Listen("tcp", ":"+GQUE_PORT)

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
