package config

type MysqlConfig struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	ShowSql  bool   `yaml:"show_sql" mapstructure:"show_sql"`
}

func (conf *MysqlConfig) HasMySql() bool {
	return conf.Host != "" && conf.Port > 0 && conf.Database != ""
}
