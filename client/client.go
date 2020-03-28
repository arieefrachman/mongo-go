package main

import (
	"context"
	"fmt"
	"github.com/arieefrachman/mongo-go/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//CreateBlogClient()
	ReadBlogClient()
}

func ReadBlogClient()  {
	s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect server: %v", err)
	}

	defer s.Close()

	c := pb.NewBlogServiceClient(s)

	blog, err := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{
		BlogId: "5e79664aa330b908c073d19d",
	})

	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	fmt.Println(blog.Blog)

}

func CreateBlogClient() {
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
