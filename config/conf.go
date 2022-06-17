package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync/atomic"
)

// DefaultConf ...
type DefaultConf struct {
	DB struct {
		Mysql struct {
			Config struct {
				DSN string `yaml:"dsn"`
			} `yaml:"config"`
		} `yaml:"mysql"`
	} `yaml:"db"`
	LoadDBInterval float64 `yaml:"load_db_interval"`
}

// ConfCache ...
var ConfCache atomic.Value

// GetConf 获取最新的默认配置
func GetConf() *DefaultConf {
	return ConfCache.Load().(*DefaultConf)
}

// LoadConf 加载并解析 Conf 配置
func LoadConf() error {
	var c *os.File
	var err error
	runInLocal := os.Getenv("GO_ENV") == "local"
	if !runInLocal {
		c, err = os.Open("default.yaml")
	} else {
		// TODO: 远程拉取配置
	}
	if err != nil {
		log.Fatalf("LoadConf fail: %v", err)
		return err
	}
	conf := &DefaultConf{}
	err = yaml.NewDecoder(c).Decode(conf)
	if err != nil {
		log.Fatalf("LoadConf fail: %v", err)
		return err
	}
	ConfCache.Store(conf)
	if !runInLocal {
		// TODO：监听器，开goroutine刷新配置
	}
	return nil
}
