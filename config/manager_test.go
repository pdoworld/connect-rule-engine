package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigManager(t *testing.T) {
	// 创建临时文件作为配置文件
	tempFile, err := os.CreateTemp("", "config_test")
	if err != nil {
		t.Fatalf("无法创建临时文件: %v", err)
	}
	defer os.Remove(tempFile.Name())

	cm := &ConfigManager{
		Configs:    make(map[string][]byte),
		ConfigPath: tempFile.Name(),
	}

	// 测试 UpdateConfig
	instanceID := "test_instance"
	testConfig := []byte("test_config_data")

	err = cm.UpdateConfig(instanceID, testConfig)
	assert.NoError(t, err)

	// 验证配置是否已保存到内存
	assert.Equal(t, testConfig, cm.Configs[instanceID])

	// 验证配置是否已写入文件
	fileContent, err := os.ReadFile(tempFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, testConfig, fileContent)

	// 测试 LoadConfig
	loadedConfig, err := cm.LoadConfig(instanceID)
	assert.NoError(t, err)
	assert.Equal(t, testConfig, loadedConfig)

	// 测试加载不存在的配置
	_, err = cm.LoadConfig("non_existent")
	assert.Error(t, err)
}
