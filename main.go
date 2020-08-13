package main

import (
	"fmt"
	grpc2 "github.com/yigitsadic/onetimecode/grpc/grpc"
	"github.com/yigitsadic/onetimecode/models"
	"github.com/yigitsadic/onetimecode/otcgo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	expiration := 60
	codeStore := models.NewCodeStore(expiration)

	s := otcgo.Server{CodeStore: codeStore}
	grpcServer := grpc.NewServer()

	grpc2.RegisterOneTimeCodeServiceServer(grpcServer, &s)

	log.Println("GRPC server up and running")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
