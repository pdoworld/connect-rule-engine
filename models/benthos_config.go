package models

import (
	"gorm.io/gorm"
)

type BenthosConfig struct {
	gorm.Model
	ConfigName string `json:"config_name"`
	Config     string `json:"config"`
}
