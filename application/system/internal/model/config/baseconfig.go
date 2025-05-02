package config

import "github.com/bingodfok/freshguard/pkg/model/config"

type BaseConfig struct {
	AppName string             `yaml:"app_name" mapstructure:"app_name"`
	Mysql   config.MysqlConfig `yaml:"mysql" mapstructure:"mysql"`
	Redis   config.RedisConfig `yaml:"redis" mapstructure:"redis"`
	Nacos   config.NacosConfig `yaml:"nacos" mapstructure:"nacos"`
}
