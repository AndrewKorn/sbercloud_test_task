package models

import (
	"github.com/joeshaw/json-lossless"
	"gorm.io/gorm"
)

type Data struct {
	lossless.JSON `json:"-"`
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Key      string `json:",any"`
	Value    string `json:"value"`
	ConfigID uint   `json:"configId"`
}

type Config struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	ServiceName string `json:"serviceName"`
	Data        []Data `gorm:"foreignKey:ConfigID" json:"data"`
	Version     uint   `json:"version"`
}
