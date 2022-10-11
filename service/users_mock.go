package service

import (
	"context"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m MockUserRepo) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockUserRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockUserRepo) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(models.User), args.Error(1)
}
