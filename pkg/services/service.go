package services

import "sbercloud_test/pkg/models"

type Service interface {
	CreateServiceConfig(serviceName string, data []models.Data) *models.Config
	GetServiceConfig(serviceName string) []models.Data
	GetServiceConfigByVersion(serviceName string, version uint) []models.Data
	DeleteServiceConfig(serviceName string) *models.Config
	DeleteServiceConfigByVersion(serviceName string, version uint) *models.Config
}
