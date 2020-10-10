package service

import (
	"IDE/GolangCloud/net.mcsql/config"
	"IDE/GolangCloud/net.mcsql/discover"
	"context"
	"errors"
)

type Service interface {

	//健康检查接口
	HealthCheck() bool

	//发送请接口
	SayHello() string

	//服务发现接口
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}

type DiscoverServiceImpl struct {
	discoveryClient discover.DiscoveryClient
}

func (service *DiscoverServiceImpl) DiscoveryService(ctx context.Context, serviceNmae string) ([]interface{}, error) {
	//从 consul 中各具服务名获取实例列表
	services := service.discoveryClient.DiscoveryServices(serviceNmae, config.Logger)

	if services == nil || len(services) == 0 {
		return nil, errors.New("instances are not existed")
	}
	return services, nil
}
