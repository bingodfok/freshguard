package alc

import (
	"fmt"
	"github.com/bingodfok/freshguard/application/resources/internal/model/config"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/region"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/spf13/viper"
)

type ApplicationContext struct {
	Config   *config.BaseConfig
	Uploader *uploader.UploadManager
}

func NewApplicationContext() (*ApplicationContext, error) {
	// 加载配置文件
	v := viper.New()
	v.SetConfigFile("./application/resources/etc/application.yaml")
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
	// 构建七牛云文件uploader
	qiniu := baseConfig.Qiniu
	manager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: credentials.NewCredentials(qiniu.AccessKey, qiniu.SecretKey),
			Regions:     region.GetRegionByID("z2", true),
		},
	})
	return &ApplicationContext{
		Config:   baseConfig,
		Uploader: manager,
	}, nil
}
