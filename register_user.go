package main

import (
	"context"
	"github.com/zvash/grpc-client/internal/authpb"
	"log"
	"time"
)

type RegisterUserArgs struct {
	Email    string
	Name     string
	Password string
}

func (a RegisterUserArgs) GetMap() map[string]string {
	args := make(map[string]string)
	args["email"] = a.Email
	args["password"] = a.Password
	args["name"] = a.Name
	return args
}

func (a RegisterUserArgs) Invoke() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	authClient := buildAuthClient()

	registerResponse, err := authClient.RegisterUser(ctx, &authpb.RegisterUserRequest{
		Email:                a.Email,
		Password:             a.Password,
		Name:                 a.Name,
		PasswordConfirmation: a.Password,
	})
	if err != nil {
		log.Panicf("Failed to register the new user, err: %v", err)
	}
	user := registerResponse.GetUser()
	log.Printf("id: %s, email: %s, name: %s", user.GetName(), user.GetEmail(), user.GetName())
}
