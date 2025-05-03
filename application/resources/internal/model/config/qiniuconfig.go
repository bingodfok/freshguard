package config

type QiniuConfig struct {
	AccessKey     string `yaml:"access_key" mapstructure:"access_key"`
	SecretKey     string `yaml:"secret_key" mapstructure:"secret_key"`
	PublicBucket  string `yaml:"public_bucket" mapstructure:"public_bucket"`
	PrivateBucket string `yaml:"private_bucket" mapstructure:"private_bucket"`
	DomainUrl     string `yaml:"domain_url" mapstructure:"domain_url"`
}
