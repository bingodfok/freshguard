package config

type WebConfig struct {
	Port        int    `json:"port" mapstructure:"port"`
	ContextPath string `json:"context_path" mapstructure:"context_path"`
}
