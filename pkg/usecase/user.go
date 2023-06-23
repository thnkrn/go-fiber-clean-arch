package usecase

import (
	"errors"

	domain "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
	iRepository "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/repository/interfaces"
	eUsecase "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/usecase/error"
	iUsecase "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/usecase/interfaces"
	"github.com/valyala/fasthttp"
)

type userUseCase struct {
	userRepo iRepository.UserRepository
}

func NewUserUseCase(repo iRepository.UserRepository) iUsecase.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) FindAll(ctx *fasthttp.RequestCtx) ([]domain.User, error) {
	users, err := c.userRepo.FindAll(ctx)
	if err == nil && len(users) == 0 {
		return users, eUsecase.NewErrorNotFound(errors.New("users not found"))
	}
	return users, err
}

func (c *userUseCase) FindByID(ctx *fasthttp.RequestCtx, id string) (domain.User, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	if err == nil && user == (domain.User{}) {
		return user, errors.New("users not found")
	}

	return user, err
}

func (c *userUseCase) Create(ctx *fasthttp.RequestCtx, user domain.User) (domain.User, error) {
	user, err := c.userRepo.Create(ctx, user)

	return user, err
}

func (c *userUseCase) Delete(ctx *fasthttp.RequestCtx, user domain.User) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}

func (c *userUseCase) UpdateByID(ctx *fasthttp.RequestCtx, id string, user domain.User) (domain.User, error) {
	user, err := c.userRepo.UpdateByID(ctx, id, user)

	return user, err
}

func (c *userUseCase) GetMatchName(ctx *fasthttp.RequestCtx, text string) ([]domain.User, error) {
	users, err := c.userRepo.GetMatchName(ctx, text)
	return users, err
}
