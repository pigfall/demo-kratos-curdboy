//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/pigfall/demo-kratos-curdboy/ent"
	"github.com/pigfall/demo-kratos-curdboy/internal/biz"
	"github.com/pigfall/demo-kratos-curdboy/internal/conf"
	"github.com/pigfall/demo-kratos-curdboy/internal/data"
	"github.com/pigfall/demo-kratos-curdboy/internal/server"
	"github.com/pigfall/demo-kratos-curdboy/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *ent.Client) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
