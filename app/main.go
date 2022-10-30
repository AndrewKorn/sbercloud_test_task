package main

import (
	"sbercloud_test/pkg/handler"
	"sbercloud_test/pkg/repositories"
	"sbercloud_test/pkg/server"
	"sbercloud_test/pkg/services"
)

func main() {
	var repository repositories.Repository
	repository = repositories.NewSQLLiteRepository()

	var service services.Service
	service = services.NewMyService(&repository)

	handler := handler.NewHandler(&service)
	server := server.NewServer(handler)
	server.Run()
}
