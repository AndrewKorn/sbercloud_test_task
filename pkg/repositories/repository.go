package repositories

import "sbercloud_test/pkg/models"

type Repository interface {
	CreateServiceConfig(config *models.Config) *models.Config
	GetServiceConfig(serviceName string) *models.Config
	GetServiceConfigByVersion(serviceName string, version uint) *models.Config
	DeleteServiceConfig(serviceName string) *models.Config
	DeleteServiceConfigByVersion(serviceName string, version uint) *models.Config
	CreateData(serviceName, key, value string) *models.Data
	GetData(serviceName string) []models.Data
	GetDataByVersion(serviceName string, version uint) []models.Data
}
