package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close()
	c := simple.NewSimpleClient(conn)
	r, err := c.SayHello(context.Background(), &simple.HelloRequest{Name: "Silence"})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(r.Message)
}