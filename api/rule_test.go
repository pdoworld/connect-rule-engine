package api

import (
	"bytes"
	"connect-rule-engine/config"
	"connect-rule-engine/models"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSetupRouter(t *testing.T) {
	// 设置测试数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("无法设置测试数据库: %v", err)
	}
	db.AutoMigrate(&models.BenthosConfig{})

	// 创建 ConfigManager
	cm := &config.ConfigManager{Configs: make(map[string][]byte)}

	// 设置路由
	router := SetupRouter(cm, db)

	// 测试 GET /rules
	t.Run("GET /rules", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/rules", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response []models.BenthosConfig
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response) // 初始应该为空
	})

	// 测试 GET /rules/:config_name
	t.Run("GET /rules/:config_name", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/rules/test_config", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
		
	// 测试 DELETE /rules/:config_name
	t.Run("DELETE /rules/:config_name", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/rules/test_config", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试 POST /rules
	t.Run("POST /rules", func(t *testing.T) {
		newConfig := models.BenthosConfig{
			ConfigName: "test_config",
			Config:     "input:\n  type: http_server\n\noutput:\n  type: stdout",
		}
		jsonData, _ := json.Marshal(newConfig)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/rules", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// 验证配置是否已保存到数据库
		var savedConfig models.BenthosConfig
		db.Where("config_name = ?", newConfig.ConfigName).First(&savedConfig)
		assert.Equal(t, newConfig.ConfigName, savedConfig.ConfigName)
		assert.Equal(t, newConfig.Config, savedConfig.Config)
	})
}

func TestStartBenthosInstance(t *testing.T) {
	ctx := context.Background()
	instanceID := "test_instance"
	config := []byte("input:\n  type: generate\n  generate:\n    count: 1\n    interval: \"\"\n    mapping: 'root = {\"message\": \"test\"}'")

	err := StartBenthosInstance(ctx, instanceID, config)
	assert.NoError(t, err)

	// 注意：这里我们只测试了函数是否返回错误
	// 实际上，我们可能需要更复杂的测试来验证 Benthos 实例是否正确启动和运行
}
