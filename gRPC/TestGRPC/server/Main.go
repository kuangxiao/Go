package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const(
	port = ":50051"
)

type Server struct{ }

func (s *Server) SayHello(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	fmt.Println("From client <<<"+req.Name)

	return &simple.HelloReplay{Message:"hello >>>" + req.Name},nil
}

func main(){
	lis,err := net.Listen("tcp",port)
	if err != nil {
		log.Fatal("Fail to listen! ",err)
	}

	s := grpc.NewServer()
	simple.RegisterSimpleServer(s,&Server{})
	reflection.Register(s)

	if err:= s.Serve(lis);err != nil{
		log.Fatal("Fail to server! ",err)
	}
}