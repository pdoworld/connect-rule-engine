package config

import (
	"fmt"
	"os"
)

type ConfigManager struct {
	Configs    map[string][]byte
	ConfigPath string
}

func (cm *ConfigManager) LoadConfig(instanceID string) ([]byte, error) {
	config, exists := cm.Configs[instanceID]
	if !exists {
		return nil, fmt.Errorf("configuration for instance %s not found", instanceID)
	}
	return config, nil
}

func (cm *ConfigManager) UpdateConfig(instanceID string, newConfig []byte) error {
	cm.Configs[instanceID] = newConfig
	return os.WriteFile(cm.ConfigPath, newConfig, 0644)
}
