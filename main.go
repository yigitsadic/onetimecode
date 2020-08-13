package main

import (
	"fmt"
	grpc2 "github.com/yigitsadic/onetimecode/grpc/grpc"
	"github.com/yigitsadic/onetimecode/models"
	"github.com/yigitsadic/onetimecode/otcgo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	codeExp := os.Getenv("CODE_EXP")
	var expiration int64
	if codeExp == "" {
		expiration = 60
	} else {
		expiration, err = strconv.ParseInt(codeExp, 10, 64)
	}

	codeStore := models.NewCodeStore(int(expiration))

	s := otcgo.Server{CodeStore: codeStore}
	grpcServer := grpc.NewServer()

	grpc2.RegisterOneTimeCodeServiceServer(grpcServer, &s)

	log.Printf("GRPC server up and running with %d seconds of expiration\n", expiration)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
