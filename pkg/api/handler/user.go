package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

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
		return c.JSON(NewUsesrResponse(users))
	}
}

func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	paramsId := c.Params("id")

	user, err := h.userUseCase.FindByID(c.Context(), paramsId)
	if err != nil {
		return err
	} else {
		return c.JSON(NewUserResponse(user))
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var request UserRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	userData := domain.NewUser(uuid.New(), request.Name, request.Email)
	user, err := h.userUseCase.Create(c.Context(), userData)
	if err != nil {
		return err
	} else {
		return c.JSON(NewUserResponse(user))
	}
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	user, err := h.userUseCase.FindByID(c.Context(), paramsId)
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
	var request UserRequest

	paramsId := c.Params("id")

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	_, err := h.userUseCase.FindByID(c.Context(), paramsId)
	if err != nil {
		return err
	}

	userData := domain.NewUser(uuid.MustParse(paramsId), request.Name, request.Email)

	user, err := h.userUseCase.UpdateByID(c.Context(), paramsId, userData)
	if err != nil {
		return err
	} else {
		return c.JSON(NewUserResponse(user))
	}
}

func (h *UserHandler) FindByMatchName(c *fiber.Ctx) error {
	paramsText := c.Params("text")
	users, err := h.userUseCase.GetMatchName(c.Context(), paramsText)

	if err != nil {
		return err
	} else {
		return c.JSON(NewUsesrResponse(users))
	}
}
