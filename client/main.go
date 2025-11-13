package main

import (
	pb "github.com/karthikbhandary2/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Karthik", "Bhanu", "Chetan", "Bhargav", "Sravan", "Imran"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}