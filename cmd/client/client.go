package main

import (
	"context"
	"fmt"
	"io"
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
	AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	user, err := client.AddUser(context.Background(), &pb.User{
		Id:    "999",
		Name:  "Joy",
		Email: "joy@teste.com",
	})
	if err != nil {
		fmt.Errorf("Something was wrong", err)
		return
	}
	fmt.Println(user)
}

func AddUserVerbose(client pb.UserServiceClient) {
	response, err := client.AddUserVerbose(context.Background(), &pb.User{
		Id:    "999",
		Name:  "Joy",
		Email: "joy@teste.com",
	})
	if err != nil {
		fmt.Errorf("Error on AddUserVerbose stream", err)
	}
	for {
		stream, err := response.Recv()
		if err == io.EOF { //quando a stream finalizar esse "erro" é disparado
			break
		}
		fmt.Println(stream)
	}
}
