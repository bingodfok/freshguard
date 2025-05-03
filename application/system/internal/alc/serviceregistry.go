package alc

import (
	"dubbo.apache.org/dubbo-go/v3"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"github.com/bingodfok/freshguard/pkg/model/config"
	"strconv"
)

// ServiceRegistry 服务注册
type ServiceRegistry struct {
	appName string
	nacos   *config.NacosConfig
}

func NewServiceRegistry(appName string, nacos *config.NacosConfig) *ServiceRegistry {
	return &ServiceRegistry{appName: appName, nacos: nacos}
}

// RegisterService 服务注册
func (sr *ServiceRegistry) RegisterService() error {
	instance, err := dubbo.NewInstance(
		dubbo.WithName(sr.appName),
		dubbo.WithRegistry(
			registry.WithNacos(),
			registry.WithAddress(sr.nacos.Host+":"+strconv.Itoa(sr.nacos.Port)),
			registry.WithGroup(sr.nacos.GroupId),
			registry.WithNamespace(sr.nacos.Namespace),
			registry.WithUsername(sr.nacos.Username),
			registry.WithPassword(sr.nacos.Password),
		),
		dubbo.WithProtocol(
			protocol.WithTriple(),
			protocol.WithPort(2000),
		),
	)
	if err != nil {
		return err
	}
	server, err := instance.NewServer()
	if err != nil {
		return err
	}
	err = server.Serve()
	if err != nil {
		return err
	}
	return nil
}
