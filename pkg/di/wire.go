//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/thnkrn/go-fiber-clean-arch/pkg/api"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/config"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/driver/db"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, UserSet, LogSet, HTTPSet)

	return &api.ServerHTTP{}, nil
}
