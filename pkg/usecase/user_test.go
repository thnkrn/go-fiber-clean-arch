package usecase_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
	mRepository "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/mocks/repository"
	usecase "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/usecase"
	iUsecase "github.com/thnkrn/go-fiber-crud-clean-arch/pkg/usecase/interfaces"
)

type userDependencies struct {
	mockUserRepo *mRepository.UserRepository
}

func createUserlUsecase() (iUsecase.UserUseCase, *userDependencies) {
	mockUserRepo := new(mRepository.UserRepository)
	usecase := usecase.NewUserUseCase(mockUserRepo)

	return usecase, &userDependencies{mockUserRepo}
}

func TestUserFindAll(t *testing.T) {
	t.Run("It should return user list", func(t *testing.T) {
		mockResponse := []domain.User{
			{
				ID:    uuid.New(),
				Name:  "name",
				Email: "email",
			},
		}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("FindAll", mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.FindAll(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return error if user list len = 0", func(t *testing.T) {
		mockResponse := []domain.User{}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("FindAll", mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.FindAll(context.TODO())

		assert.Equal(t, res, mockResponse)
		assert.Equal(t, "users not found", err.Error())

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserFindByID(t *testing.T) {
	t.Run("It should return user list", func(t *testing.T) {
		mockResponse := domain.User{
			ID:    uuid.New(),
			Name:  "name",
			Email: "email",
		}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("FindByID", mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.FindByID(context.TODO(), mock.Anything)

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return error if user not found", func(t *testing.T) {
		mockResponse := domain.User{}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("FindByID", mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.FindByID(context.TODO(), mock.Anything)

		assert.Equal(t, res, mockResponse)
		assert.Equal(t, "users not found", err.Error())

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserCreate(t *testing.T) {
	t.Run("It should create user success", func(t *testing.T) {
		mockResponse := domain.User{
			ID:    uuid.New(),
			Name:  "name",
			Email: "email",
		}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("Create", mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.Create(context.TODO(), mockResponse)

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserDelete(t *testing.T) {
	t.Run("It should delete user success", func(t *testing.T) {
		mockResponse := domain.User{
			ID:    uuid.New(),
			Name:  "name",
			Email: "email",
		}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := usecase.Delete(context.TODO(), mockResponse)

		assert.NoError(t, err)

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserUpdateByID(t *testing.T) {
	t.Run("It should update user success", func(t *testing.T) {
		mockResponse := domain.User{
			ID:    uuid.New(),
			Name:  "name",
			Email: "email",
		}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.UpdateByID(context.TODO(), mock.Anything, mockResponse)

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserGetMatchName(t *testing.T) {
	t.Run("It should get users by matched name success", func(t *testing.T) {
		mockResponse := []domain.User{{
			ID:    uuid.New(),
			Name:  "name",
			Email: "email",
		},
		}
		usecase, deps := createUserlUsecase()
		deps.mockUserRepo.On("GetMatchName", mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.GetMatchName(context.TODO(), mock.Anything)

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})
}
