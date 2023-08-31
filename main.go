package main

import (
	"flag"
	"github.com/zvash/grpc-client/internal/authpb"
	"github.com/zvash/grpc-client/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Invokable interface {
	GetMap() map[string]string
	Invoke()
}

func prepareMethodCall() Invokable {
	method := flag.String("method", "", "gRPC method name")
	email := flag.String("email", "", "user's email for authentication")
	name := flag.String("name", "", "user's name for update")
	password := flag.String("password", "", "user's password for authentication")
	imagePath := flag.String("image", "", "path to an image")

	flag.Parse()

	switch *method {
	case "UpdateProfile":
		return UpdateProfileArgs{
			Email:     *email,
			Name:      *name,
			Password:  *password,
			ImagePath: *imagePath,
		}
	case "RegisterUser":
		return RegisterUserArgs{
			Email:    *email,
			Name:     *name,
			Password: *password,
		}
	}
	return nil
}

func buildAuthClient() authpb.AuthClient {
	serverAddress := "127.0.0.1:9090"
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(serverAddress, transportOption)
	if err != nil {
		log.Panicf("Could not dial %s, err: %v", serverAddress, transportOption)
	}
	return authpb.NewAuthClient(cc)
}

func buildAppClient() pb.AppClient {
	serverAddress := "127.0.0.1:9090"
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(serverAddress, transportOption)
	if err != nil {
		log.Panicf("Could not dial %s, err: %v", serverAddress, transportOption)
	}
	return pb.NewAppClient(cc)
}

func main() {
	method := prepareMethodCall()
	if method == nil {
		log.Panic("invalid method.")
	}
	method.Invoke()
}
