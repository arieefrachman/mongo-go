package main

import (
	"fmt"
	"github.com/arieefrachman/mongo-go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {}
func main()  {
	fmt.Println("Blog service started...")
	listener, errLis := net.Listen("tcp", "0.0.0.0:50051")

	if errLis != nil {
		log.Fatalf("failed to listen: %v\n", errLis)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("server is starting....")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("stopping the server")
	s.Stop()
	fmt.Println("closing the listener")
	listener.Close()

}
