package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/zvash/grpc-client/internal/authpb"
	"github.com/zvash/grpc-client/internal/pb"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type UpdateProfileArgs struct {
	Email     string
	Password  string
	Name      string
	ImagePath string
}

func (a UpdateProfileArgs) GetMap() map[string]string {
	args := make(map[string]string)
	args["email"] = a.Email
	args["password"] = a.Password
	args["name"] = a.Name
	args["image_path"] = a.ImagePath
	return args
}

func (a UpdateProfileArgs) Invoke() {
	file, err := os.Open(a.ImagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	authClient := buildAuthClient()

	loginResponse, err := authClient.Login(ctx, &authpb.LoginRequest{
		Email:    a.Email,
		Password: a.Password,
	})
	if err != nil {
		log.Panicf("Failed to login, err: %v", err)
	}

	token := fmt.Sprintf("Bearer %s", loginResponse.AccessToken)

	headers := make(map[string]string)
	headers["authorization"] = token
	headers["user-agent"] = "custom-grpc-client"
	md := metadata.New(headers)
	requestContext := metadata.NewOutgoingContext(ctx, md)

	gatewayClient := buildAppClient()

	stream, err := gatewayClient.UpdateProfile(requestContext)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	extension := filepath.Ext(a.ImagePath)

	req := &pb.UpdateProfileRequest{
		Data: &pb.UpdateProfileRequest_Info{
			Info: &pb.ProfileInfo{
				Name:     &a.Name,
				ImageExt: &extension,
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send user info to server: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024*128)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.UpdateProfileRequest{
			Data: &pb.UpdateProfileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("profile updated with messege: %s", res.Message)
}
