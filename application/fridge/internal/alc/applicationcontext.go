package alc

import (
	"fmt"
	"github.com/bingodfok/freshguard/application/fridge/internal/model/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

type ApplicationContext struct {
	Config *config.BaseConfig
	DB     *xorm.Engine
}

func NewApplicationContext() (*ApplicationContext, error) {
	// 加载配置文件
	v := viper.New()
	v.SetConfigFile("./application/fridge/etc/application.yaml")
	v.SetConfigType("yaml")
	baseConfig := &config.BaseConfig{}
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("read config file failed, err:%v\n", err)
		return nil, err
	}
	if err := v.Unmarshal(baseConfig); err != nil {
		fmt.Printf("config file unmarshal fail,err:%v\n", err)
		return nil, err
	}
	// 连接MySQL
	var db *xorm.Engine
	mysqlConfig := baseConfig.Mysql
	if mysqlConfig.HasMySql() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlConfig.Username,
			mysqlConfig.Password,
			mysqlConfig.Host,
			mysqlConfig.Port,
			mysqlConfig.Database,
		)
		engine, err := xorm.NewEngine("mysql", dsn)
		if err != nil || engine.Ping() != nil {
			fmt.Printf("connect mysql failed, err:%v\n", err)
			return nil, err
		}
		if mysqlConfig.ShowSql {
			engine.ShowSQL(true)
		}
		db = engine
	}
	return &ApplicationContext{
		Config: baseConfig,
		DB:     db,
	}, nil
}

func (ctx *ApplicationContext) Close() error {
	if err := ctx.DB.Close(); err != nil {
		return err
	}
	return nil
}
