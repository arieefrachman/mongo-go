package main

import (
	"context"
	"fmt"
	"github.com/arieefrachman/mongo-go/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client started")

	s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer s.Close()

	c := pb.NewBlogServiceClient(s)

	request := &pb.Blog{
		AuthorId:             "Peram",
		Title:                "Jaman hamil dia nih",
		Content:              "sampe eneg",
	}
	createdBlog, err := c.CreateBlog(context.Background(), &pb.CreateBlogRequest{
		Blog: request,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("blog has been created: %v", createdBlog)
}
