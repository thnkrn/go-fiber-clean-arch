//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/thnkrn/go-fiber-clean-arch/pkg/api"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/api/handler"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/config"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/db"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/repository"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, HTTPSet)

	return &api.ServerHTTP{}, nil
}
