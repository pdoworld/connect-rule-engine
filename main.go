package main

import (
	"connect-rule-engine/api"
	"connect-rule-engine/config"
	"context"
	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"log"
)

func main() {
	cm := &config.ConfigManager{Configs: make(map[string][]byte)}

	// Example: loading and starting Benthos instances
	for instanceID, config := range cm.Configs {
		if err := api.StartBenthosInstance(context.Background(), instanceID, config); err != nil {
			log.Fatalf("Failed to start Benthos instance %s: %v", instanceID, err)
		}
	}

	// 启动 API 服务
	r := api.SetupRouter(cm)
	r.Run(":8080")
}
