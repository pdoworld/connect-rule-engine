package main

import (
	"connect-rule-engine/api"
	"connect-rule-engine/config"
	"context"
	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"log"
	"connect-rule-engine/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cm := &config.ConfigManager{Configs: make(map[string][]byte)}

	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("benthos_configs.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移表结构
	db.AutoMigrate(&models.BenthosConfig{})
	
	// 从数据库加载所有 Benthos 配置
	var configs []models.BenthosConfig
	if err := db.Find(&configs).Error; err != nil {
		log.Fatalf("Failed to load configurations from database: %v", err)
	}

	// 将配置加载到 ConfigManager 并启动相应的 Benthos 实例
	for _, config := range configs {
		cm.Configs[config.ConfigName] = []byte(config.Config)
		if err := api.StartBenthosInstance(context.Background(), config.ConfigName, []byte(config.Config)); err != nil {
			log.Fatalf("Failed to start Benthos instance %s: %v", config.ConfigName, err)
		}
	}

	// 启动 API 服务
	r := api.SetupRouter(cm, db)
	r.Run(":8080")
}
