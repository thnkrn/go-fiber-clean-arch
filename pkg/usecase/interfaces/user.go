package interfaces

import (
	domain "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
	"github.com/valyala/fasthttp"
)

type UserUseCase interface {
	FindAll(ctx *fasthttp.RequestCtx) ([]domain.User, error)
	FindByID(ctx *fasthttp.RequestCtx, id string) (domain.User, error)
	Create(ctx *fasthttp.RequestCtx, user domain.User) (domain.User, error)
	Delete(ctx *fasthttp.RequestCtx, user domain.User) error
	UpdateByID(ctx *fasthttp.RequestCtx, id string, user domain.User) (domain.User, error)
	GetMatchName(ctx *fasthttp.RequestCtx, text string) ([]domain.User, error)
}
