package config

type RedisConfig struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     uint   `yaml:"port" mapstructure:"port"`
	Password string `yaml:"password" mapstructure:"password"`
	Database int    `yaml:"database" mapstructure:"database"`
}
