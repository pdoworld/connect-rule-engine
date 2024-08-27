package api

import (
	"connect-rule-engine/config"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/gin-gonic/gin"

	// "os"
	"connect-rule-engine/models"

	"gorm.io/gorm"
)

func SetupRouter(cm *config.ConfigManager, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/rules", func(c *gin.Context) {
		var configs []models.BenthosConfig
		db.Find(&configs)
		c.JSON(http.StatusOK, configs)
	})

	r.POST("/rules", func(c *gin.Context) {
		var newConfig models.BenthosConfig
		if err := c.BindJSON(&newConfig); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}
	
		// 保存配置到数据库
		if err := db.Create(&newConfig).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save configuration"})
			return
		}

		// 将新配置加载到 ConfigManager 并启动 Benthos 实例
		cm.Configs[newConfig.ConfigName] = []byte(newConfig.Config)
		if err := StartBenthosInstance(context.Background(), newConfig.ConfigName, []byte(newConfig.Config)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start Benthos instance"})
			return
		}
		c.Status(http.StatusOK)
	})

	return r
}

func StartBenthosInstance(ctx context.Context, instanceID string, config []byte) error {
	// 创建 StreamBuilder
	builder := service.NewStreamBuilder()

	// 加载配置
	if err := builder.SetYAML(string(config)); err != nil {
		return fmt.Errorf("failed to set YAML config for instance %s: %w", instanceID, err)
	}

	// 构建并启动流
	stream, err := builder.Build()
	if err != nil {
		return fmt.Errorf("failed to build stream for instance %s: %w", instanceID, err)
	}

	go func() {
		if err := stream.Run(ctx); err != nil {
			log.Printf("Benthos instance %s stopped: %v", instanceID, err)
		}
	}()

	return nil
}
