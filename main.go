package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/zvash/grpc-client/internal/authpb"
	"github.com/zvash/grpc-client/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	uploadImage("email", "1234567", "internal/tmp/android-stub.jpeg", "SHN")
}

func uploadImage(username, password, imagePath, newName string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAddress := "127.0.0.1:9090"
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(serverAddress, transportOption)
	if err != nil {
		log.Panicf("Could not dial %s, err: %v", serverAddress, transportOption)
	}

	authClient := authpb.NewAuthClient(cc)

	loginResponse, err := authClient.Login(ctx, &authpb.LoginRequest{
		Email:    username,
		Password: password,
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

	gatewayClient := pb.NewAppClient(cc)

	stream, err := gatewayClient.UpdateProfile(requestContext)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	extension := filepath.Ext(imagePath)

	req := &pb.UpdateProfileRequest{
		Data: &pb.UpdateProfileRequest_Info{
			Info: &pb.ProfileInfo{
				Name:     &newName,
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
