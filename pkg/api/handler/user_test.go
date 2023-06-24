package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	handler "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/api/handler"
	domain "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
	musecase "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/mocks/usecase"
	test "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/test"
)

type userDependencies struct {
	mockUserUsecase *musecase.UserUseCase
}

func createUserHandler() (*handler.UserHandler, *userDependencies) {
	mockUserUsecase := new(musecase.UserUseCase)
	handler := handler.NewUserHandler(mockUserUsecase)

	return handler, &userDependencies{mockUserUsecase}
}

func TestUserFindAll(t *testing.T) {
	urlPattern := "/api/users"

	t.Run("It should return status 200 if get successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/api/users",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindAll", mock.Anything).Return([]domain.User{{ID: uuid.New(), Name: "name", Email: "email"}}, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.FindAll)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/api/users",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindAll", mock.Anything).Return(nil, errors.New("an error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.FindAll)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserFindByID(t *testing.T) {
	urlPattern := "/api/users/:id"

	t.Run("It should return status 200 if get successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/api/users/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{ID: uuid.New(), Name: "name", Email: "email"}, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.FindByID)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/api/users/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{}, errors.New("an error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.FindByID)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserCreate(t *testing.T) {
	urlPattern := "/api/users"

	t.Run("It should return status 200 if create successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "POST",
			Path:   "/api/users",
			Body: handler.UserRequest{
				Name:  "name",
				Email: "email",
			},
		}

		jsonBytes, err := json.Marshal(handlerRequest.Body)
		assert.NoError(t, err)

		jsonBody := bytes.NewReader(jsonBytes)

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, jsonBody)
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Create", mock.Anything, mock.Anything).Return(domain.User{ID: uuid.New(), Name: "name", Email: "email"}, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Create)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "POST",
			Path:   "/api/users",
			Body: handler.UserRequest{
				Name:  "name",
				Email: "email",
			},
		}

		jsonBytes, err := json.Marshal(handlerRequest.Body)
		assert.NoError(t, err)

		jsonBody := bytes.NewReader(jsonBytes)

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, jsonBody)
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Create", mock.Anything, mock.Anything).Return(domain.User{ID: uuid.New(), Name: "name", Email: "email"}, errors.New("error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Create)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})

	t.Run("It should return status 500 if body cannot parsed", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "POST",
			Path:   "/api/users",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, strings.NewReader("failed"))
		assert.NoError(t, err)

		handler, deps := createUserHandler()

		response := test.RequestHandler(urlPattern, request, handler.Create)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserDelete(t *testing.T) {
	urlPattern := "/api/users/:id"

	t.Run("It should return status 200 if delete successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "DELETE",
			Path:   "/api/users/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{ID: uuid.New(), Name: "name", Email: "email"}, nil).Once()
		deps.mockUserUsecase.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Delete)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured from FindByID usecase", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "DELETE",
			Path:   "/api/users/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{ID: uuid.New(), Name: "name", Email: "email"}, errors.New("error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Delete)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured from Delete usecase", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "DELETE",
			Path:   "/api/users/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{ID: uuid.New(), Name: "name", Email: "email"}, nil).Once()
		deps.mockUserUsecase.On("Delete", mock.Anything, mock.Anything).Return(errors.New("error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Delete)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserFindByMatchName(t *testing.T) {
	urlPattern := "/api/users/:text"

	t.Run("It should return status 200 if find by matched name successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/api/users/name",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("GetMatchName", mock.Anything, mock.Anything).Return([]domain.User{{ID: uuid.New(), Name: "name", Email: "email"}}, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.FindByMatchName)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/api/users/name",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("GetMatchName", mock.Anything, mock.Anything).Return([]domain.User{{ID: uuid.New(), Name: "name", Email: "email"}}, errors.New("error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.FindByMatchName)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserUpdate(t *testing.T) {
	urlPattern := "/api/users/:id"

	t.Run("It should return status 200 if update successfully", func(t *testing.T) {
		mockUUID := uuid.New()
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/api/users/" + mockUUID.String(),
			Body: handler.UserRequest{
				Name:  "newName",
				Email: "newEmail",
			},
		}

		jsonBytes, err := json.Marshal(handlerRequest.Body)
		assert.NoError(t, err)

		jsonBody := bytes.NewReader(jsonBytes)

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, jsonBody)
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{ID: mockUUID, Name: "name", Email: "email"}, nil).Once()
		deps.mockUserUsecase.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(domain.User{ID: mockUUID, Name: "newName", Email: "newEmail"}, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Update)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured from FindByID usecase", func(t *testing.T) {
		mockUUID := uuid.New()
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/api/users/" + mockUUID.String(),
			Body: handler.UserRequest{
				Name:  "newName",
				Email: "newEmail",
			},
		}

		jsonBytes, err := json.Marshal(handlerRequest.Body)
		assert.NoError(t, err)

		jsonBody := bytes.NewReader(jsonBytes)

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, jsonBody)
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{ID: mockUUID, Name: "name", Email: "email"}, errors.New("error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Update)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})

	t.Run("It should return status 500 if error occured from UpdateByID usecase", func(t *testing.T) {
		mockUUID := uuid.New()
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/api/users/" + mockUUID.String(),
			Body: handler.UserRequest{
				Name:  "newName",
				Email: "newEmail",
			},
		}

		jsonBytes, err := json.Marshal(handlerRequest.Body)
		assert.NoError(t, err)

		jsonBody := bytes.NewReader(jsonBytes)

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, jsonBody)
		request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{ID: mockUUID, Name: "name", Email: "email"}, nil).Once()
		deps.mockUserUsecase.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(domain.User{ID: mockUUID, Name: "name", Email: "email"}, errors.New("error occured")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Update)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})

	t.Run("It should return status 500 if body cannot parsed", func(t *testing.T) {
		mockUUID := uuid.New()
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/api/users/" + mockUUID.String(),
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, strings.NewReader("failed"))
		assert.NoError(t, err)

		handler, deps := createUserHandler()

		response := test.RequestHandler(urlPattern, request, handler.Update)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}
