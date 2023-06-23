package repository

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/valyala/fasthttp"

	domain "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/repository/interfaces"
)

type User struct {
	ID    uuid.UUID `gorm:"type:UUID;primaryKey"`
	Name  string    `gorm:"not null"`
	Email string    `gorm:"not null"`
}

func NewUser(u domain.User) User {
	var user User
	copier.Copy(&user, &u)

	return user
}

func (u *User) ToUser() domain.User {
	model := domain.NewUser(u.ID, u.Name, u.Email)
	return model
}

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) FindAll(ctx *fasthttp.RequestCtx) ([]domain.User, error) {
	var pUsers []User
	tx := c.DB.Find(&pUsers)

	users := make([]domain.User, len(pUsers))
	for i, v := range pUsers {
		users[i] = v.ToUser()
	}

	return users, tx.Error
}

func (c *userDatabase) FindByID(ctx *fasthttp.RequestCtx, id string) (domain.User, error) {
	var pUser User
	tx := c.DB.Where("id = ?", id).Find(&pUser)

	return pUser.ToUser(), tx.Error
}

func (c *userDatabase) Create(ctx *fasthttp.RequestCtx, user domain.User) (domain.User, error) {
	pUser := NewUser(user)
	tx := c.DB.Create(pUser)

	return user, tx.Error
}

func (c *userDatabase) Delete(ctx *fasthttp.RequestCtx, user domain.User) error {
	pUser := NewUser(user)
	tx := c.DB.Delete(pUser)

	return tx.Error
}

func (c *userDatabase) UpdateByID(ctx *fasthttp.RequestCtx, id string, user domain.User) (domain.User, error) {
	pUser := NewUser(user)

	tx := c.DB.Model(&user).Where("id = ?", id).Updates(pUser)
	if tx.Error != nil {
		return user, tx.Error
	}

	tx = c.DB.Where("id = ?", id).Find(&pUser)

	return pUser.ToUser(), tx.Error
}

func (c *userDatabase) GetMatchName(ctx *fasthttp.RequestCtx, text string) ([]domain.User, error) {
	var pUsers []User

	name := "%" + text + "%"
	tx := c.DB.Where("name LIKE ?", name).Find(&pUsers)

	users := make([]domain.User, len(pUsers))
	for i, v := range pUsers {
		users[i] = v.ToUser()
	}

	return users, tx.Error
}
