package main

import (
	"cobasolid/repository"
	"cobasolid/service"
	"context"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	ctx := context.Background()
	repo := repository.NewRepository(nil)
	svc := service.NewUserManagementService(repo)
	svc.UserRegistration(ctx, nil)
}
