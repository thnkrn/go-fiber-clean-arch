package repository

import (
	"gorm.io/gorm"

	domain "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/repository/interfaces"
	"github.com/valyala/fasthttp"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) FindAll(ctx *fasthttp.RequestCtx) ([]domain.User, error) {
	var users []domain.User
	tx := c.DB.Find(&users)

	return users, tx.Error
}

func (c *userDatabase) FindByID(ctx *fasthttp.RequestCtx, id uint) (domain.User, error) {
	var user domain.User
	tx := c.DB.First(&user, id)

	return user, tx.Error
}

func (c *userDatabase) Create(ctx *fasthttp.RequestCtx, user domain.User) (domain.User, error) {
	tx := c.DB.Create(&user)

	return user, tx.Error
}

func (c *userDatabase) Delete(ctx *fasthttp.RequestCtx, user domain.User) error {
	tx := c.DB.Delete(&user)

	return tx.Error
}

func (c *userDatabase) UpdateByID(ctx *fasthttp.RequestCtx, id uint, user domain.User) (domain.User, error) {
	tx := c.DB.Model(&user).Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return user, tx.Error
	}

	var ruser domain.User
	tx = tx.First(&ruser, id)

	return ruser, tx.Error
}

func (c *userDatabase) GetMatchName(ctx *fasthttp.RequestCtx, text string) ([]domain.User, error) {
	var users []domain.User
	name := "%" + text + "%"
	tx := c.DB.Where("name LIKE ?", name).Find(&users)

	return users, tx.Error
}
