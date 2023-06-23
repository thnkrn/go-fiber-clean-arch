package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"

	domain "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
	iUsecase "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/usecase/interfaces"
)

type UserHandler struct {
	userUseCase iUsecase.UserUseCase
}

func NewUserHandler(usecase iUsecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (h *UserHandler) FindAll(c *fiber.Ctx) error {
	users, err := h.userUseCase.FindAll(c.Context())

	if err != nil {
		return err
	} else {
		var response []UserResponse
		copier.Copy(&response, &users)

		return c.JSON(response)
	}
}

func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return err
	}

	users, err := h.userUseCase.FindByID(c.Context(), uint(id))
	if err != nil {
		return err
	} else {
		var response UserResponse
		copier.Copy(&response, &users)

		return c.JSON(response)
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user, err := h.userUseCase.Create(c.Context(), user)
	if err != nil {
		return err
	} else {
		response := UserResponse{}
		copier.Copy(&response, &user)

		return c.JSON(response)
	}
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return err
	}

	user, err := h.userUseCase.FindByID(c.Context(), uint(id))
	if err != nil {
		return err
	}

	err = h.userUseCase.Delete(c.Context(), user)
	if err != nil {
		return err
	} else {
		return c.SendStatus(http.StatusNoContent)
	}
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	var user domain.User

	paramsId := c.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	_, err = h.userUseCase.FindByID(c.Context(), uint(id))
	if err != nil {
		return err
	}

	user, err = h.userUseCase.UpdateByID(c.Context(), uint(id), user)
	fmt.Printf("%%v: %v\n", user)
	if err != nil {
		return err
	} else {
		response := UserResponse{}
		copier.Copy(&response, &user)

		return c.JSON(response)
	}
}

func (h *UserHandler) FindByMatchName(c *fiber.Ctx) error {
	paramsText := c.Params("text")
	users, err := h.userUseCase.GetMatchName(c.Context(), paramsText)

	if err != nil {
		return err
	} else {
		var response []UserResponse
		copier.Copy(&response, &users)

		return c.JSON(response)
	}
}
