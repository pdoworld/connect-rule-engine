package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestBenthosConfig(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("无法设置测试数据库: %v", err)
	}

	// 自动迁移架构
	db.AutoMigrate(&BenthosConfig{})

	// 创建测试配置
	testConfig := BenthosConfig{
		ConfigName: "test_config",
		Config:     "input:\n  type: http_server\n\noutput:\n  type: stdout",
	}

	// 保存到数据库
	result := db.Create(&testConfig)
	assert.NoError(t, result.Error)
	assert.NotZero(t, testConfig.ID)

	// 从数据库读取
	var retrievedConfig BenthosConfig
	db.First(&retrievedConfig, testConfig.ID)

	// 验证字段
	assert.Equal(t, testConfig.ConfigName, retrievedConfig.ConfigName)
	assert.Equal(t, testConfig.Config, retrievedConfig.Config)
}
