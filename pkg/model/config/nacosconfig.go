package config

type NacosConfig struct {
	Host      string `yaml:"host" mapstructure:"host"`
	Port      int    `yaml:"port" mapstructure:"port"`
	Namespace string `yaml:"namespace" mapstructure:"namespace"`
	GroupId   string `yaml:"group_id" mapstructure:"group_id"`
	Username  string `yaml:"username" mapstructure:"username"`
	Password  string `yaml:"password" mapstructure:"password"`
}
