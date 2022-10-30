package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sbercloud_test/pkg/models"
)

type SQLLiteRepository struct {
	db *gorm.DB
}

func (r *SQLLiteRepository) CreateServiceConfig(config *models.Config) *models.Config {
	r.db.Create(&config)
	return config
}

func (r *SQLLiteRepository) GetServiceConfig(serviceName string) *models.Config {
	config := &models.Config{}
	r.db.Where(models.Config{ServiceName: serviceName}).Last(&config)
	return config
}

func (r *SQLLiteRepository) GetServiceConfigByVersion(serviceName string, version uint) *models.Config {
	config := &models.Config{}
	r.db.Where(models.Config{ServiceName: serviceName, Version: version}).First(&config)
	return config
}

func (r *SQLLiteRepository) DeleteServiceConfig(serviceName string) *models.Config {
	config := &models.Config{}
	r.db.Where(models.Config{ServiceName: serviceName}).Delete(config)
	return config
}

func (r *SQLLiteRepository) DeleteServiceConfigByVersion(serviceName string, version uint) *models.Config {
	config := &models.Config{}
	r.db.Where(models.Config{ServiceName: serviceName, Version: version}).Delete(config)
	return config
}

func (r *SQLLiteRepository) CreateData(serviceName, key, value string) *models.Data {
	config := r.GetServiceConfig(serviceName)
	data := &models.Data{Key: key, Value: value, ConfigID: config.ID}
	err := r.db.Model(&config).Association("Data").Append(data)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func (r *SQLLiteRepository) GetData(serviceName string) []models.Data {
	config := r.GetServiceConfig(serviceName)

	var data []models.Data
	err := r.db.Model(&config).Association("Data").Find(&data)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func (r *SQLLiteRepository) GetDataByVersion(serviceName string, version uint) []models.Data {
	config := r.GetServiceConfigByVersion(serviceName, version)

	var data []models.Data
	err := r.db.Model(&config).Association("Data").Find(&data)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func NewSQLLiteRepository() *SQLLiteRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Config{}, &models.Data{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &SQLLiteRepository{db: db}
}
