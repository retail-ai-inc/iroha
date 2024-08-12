package config

import (
	_ "embed"
	"encoding/json"
	"sync"
)

type Config struct {
	Algo      []string          `json:"algo"`
	Namespace map[string]Tenant `json:"namespace"`
}

type Tenant struct {
	Token        string   `json:"token"`
	Algo         string   `json:"algo"`
	Salt         string   `json:"salt"`
	WhiteListIps []string `json:"whiteListIps"`
}

var (
	configInstance *Config
	once           sync.Once
)

//go:embed env.json
var envString string

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{}
		err := json.Unmarshal([]byte(envString), configInstance)
		if err != nil {
			panic(err)
		}
	})
	return configInstance
}
