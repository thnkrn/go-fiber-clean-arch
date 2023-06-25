package di

import (
	"github.com/google/wire"

	"github.com/thnkrn/go-fiber-clean-arch/pkg/api/handler"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/repository"
	"github.com/thnkrn/go-fiber-clean-arch/pkg/usecase"
)

var UserSet = wire.NewSet(
	repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler,
)
