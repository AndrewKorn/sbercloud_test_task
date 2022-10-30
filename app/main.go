package main

import (
	"sbercloud_test/pkg/handler"
	"sbercloud_test/pkg/repositories"
	"sbercloud_test/pkg/server"
	"sbercloud_test/pkg/services"
)

func main() {
	var repo repositories.Repository
	repo = repositories.NewSQLLiteRepository()

	var service services.Service
	service = services.NewMyService(&repo)

	h := handler.NewHandler(&service)
	server := server.NewServer(h)
	server.Run()
}
