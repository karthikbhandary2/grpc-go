package main

import (
	"context"
	"log"
	"time"

	pb "github.com/karthikbhandary2/go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		panic(err)
	}
	log.Printf("%s", res.Message)
}