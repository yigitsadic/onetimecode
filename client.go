package main

import (
	grpc2 "github.com/yigitsadic/onetimecode/grpc/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := grpc2.NewOneTimeCodeServiceClient(conn)

	identifiers := []string{
		"0b8f79ab-4cd9-4071-8f5e-a9c02c11b3c6",
		"c392c7fd-437f-48b0-948e-7cc36a2d893e",
		"d274f465-0ee2-49c6-9915-572cf62dc5b1",
		"41cb9848-1de2-4f89-bfbb-79269bd93594",
		"bc05216c-a8f9-4b29-a829-2a2ab0d3c3c4",
		"2284079e-c88e-40ff-8a16-bea09b81493d",
	}

	var resp *grpc2.OneTimeCodeResponse

	// Create example codes
	for _, id := range identifiers {
		var err error

		resp, err = c.CreateCode(context.Background(), &grpc2.OneTimeCodeGen{Identifier: id})
		if err != nil {
			log.Fatalf("Error during creation of code. %s\n", err)
		}

		log.Printf("Code created:\t\tIDENTIFIER=%s\tCODE=%s\n", id, resp.Value)

		resp, err := c.ReadCode(context.Background(), &grpc2.ReadCodeReq{Value: resp.Value})
		if err != nil {
			log.Fatalf("Error during reading of code. %s - %s\n", resp.Value, err)
		}

		log.Printf("Code readed:\t\tIDENTIFIER:%s\tCODE=%s\n\n", id, resp.Value)
	}
}
