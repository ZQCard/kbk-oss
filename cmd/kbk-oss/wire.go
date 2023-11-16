//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"repo.example.com/oss/internal/biz"
	"repo.example.com/oss/internal/conf"
	"repo.example.com/oss/internal/data"
	"repo.example.com/oss/internal/server"
	"repo.example.com/oss/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *conf.Server, *conf.Data, *conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
