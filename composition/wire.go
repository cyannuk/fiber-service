//+build wireinject

package composition

import (
	"fiber-service/api"
	"fiber-service/config"
	"fiber-service/domain/repository"
	"fiber-service/domain/service"
	"github.com/google/wire"
)

func Application() (api.Application, error) {
	wire.Build(config.GetDatabaseConfig, config.GetServerConfig, repository.NewDataSource, repository.NewUserRepository, repository.NewOrderRepository,
		service.NewUserService, service.NewOrderService, api.NewApplication)
	return nil, nil
}
