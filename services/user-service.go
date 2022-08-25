package services

import (
	"context"
	"time"

	"github.com/leonardodelira/go-grpc/pb"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() pb.UserServiceServer {
	return &userService{}
}

func (*userService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User{
		Id:    "1",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*userService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User inserted",
		User: &pb.User{
			Id: "9999",
		},
	})

	return nil
}
