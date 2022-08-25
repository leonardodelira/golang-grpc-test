package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	AddUsers(client)
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

func AddUsers(client pb.UserServiceClient) {
	users := []*pb.User{
		{
			Id:    "1",
			Name:  "name 1",
			Email: "email1@teste",
		},
		{
			Id:    "2",
			Name:  "name 2",
			Email: "email2@teste",
		},
		{
			Id:    "3",
			Name:  "name 3",
			Email: "email3@teste",
		},
	}

	stream, _ := client.AddUsers(context.Background())

	for _, user := range users {
		stream.Send(user)
		time.Sleep(time.Second * 2)
	}

	response, _ := stream.CloseAndRecv()
	fmt.Println(response)
}
