package api

import (
	"connect-rule-engine/config"
	"context"
	"fmt"
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func SetupRouter(cm *config.ConfigManager) *gin.Engine {
	r := gin.Default()

	r.GET("/rules/:instanceID", func(c *gin.Context) {
		instanceID := c.Param("instanceID")
		config, err := cm.LoadConfig(instanceID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"config": string(config)})
	})

	r.POST("/rules/:instanceID", func(c *gin.Context) {
		instanceID := c.Param("instanceID")
		var newConfig []byte
		// if err := c.BindJSON(&newConfig); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// if err := cm.UpdateConfig(instanceID, newConfig); err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }
		newConfig, err := os.ReadFile("/Users/yangguang/code/toys/test/connect-rule-engine/benthos.yaml")
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}
		log.Println("Configuration loaded successfully:")
		log.Println(string(newConfig))
		// 重新启动对应的 Benthos 实例
		go func() {
			if err := StartBenthosInstance(context.Background(), instanceID, newConfig); err != nil {
				log.Fatalf("Failed to restart Benthos instance %s: %v", instanceID, err)
			}
		}()

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
