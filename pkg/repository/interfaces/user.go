package interfaces

import (
	"context"

	domain "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id uint) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
	UpdateByID(ctx context.Context, id uint, user domain.User) (domain.User, error)
	GetMatchName(ctx context.Context, text string) ([]domain.User, error)
}
