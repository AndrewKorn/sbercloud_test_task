package services

import (
	"sbercloud_test/pkg/models"
	"sbercloud_test/pkg/repositories"
)

type MyService struct {
	repos *repositories.Repository
}

func (service *MyService) CreateServiceConfig(serviceName string, data []models.Data) *models.Config {
	r := *service.repos

	lastVersion := 0
	last := r.GetServiceConfig(serviceName)
	if last != nil {
		lastVersion = int(last.Version)
	}

	config := &models.Config{ServiceName: serviceName, Version: uint(lastVersion) + 1}
	r.CreateServiceConfig(config)
	for i := range data {
		r.CreateData(serviceName, data[i].Key, data[i].Value)
	}

	return config
}

func (service *MyService) GetServiceConfig(serviceName string) []models.Data {
	r := *service.repos
	data := r.GetData(serviceName)
	return data
}

func (service *MyService) GetServiceConfigByVersion(serviceName string, version uint) []models.Data {
	r := *service.repos
	data := r.GetDataByVersion(serviceName, version)
	return data
}

func (service *MyService) DeleteServiceConfig(serviceName string) *models.Config {
	r := *service.repos
	deleted := r.DeleteServiceConfig(serviceName)
	return deleted
}

func (service *MyService) DeleteServiceConfigByVersion(serviceName string, version uint) *models.Config {
	r := *service.repos
	deleted := r.DeleteServiceConfigByVersion(serviceName, version)
	return deleted
}

func NewMyService(repository *repositories.Repository) *MyService {
	return &MyService{repos: repository}
}
