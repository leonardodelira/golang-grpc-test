package main

import (
	"context"
	"fmt"
	"log"

	"github.com/leonardodelira/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error on dial server %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	user, err := client.AddUser(context.Background(), &pb.User{
		Id:    "999",
		Name:  "Joy",
		Email: "joy@teste.com",
	})
	if err != nil {
		fmt.Errorf("Something was wrong", err)
		return
	}
	fmt.Print(user)
}
