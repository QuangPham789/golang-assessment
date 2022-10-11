package repository

import (
	"context"
	models "github.com/quangpham789/golang-assessment/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	GetUserByID(ctx context.Context, id int) (*models.User, error)

	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type RelationsRepo interface {
	CreateRelation(ctx context.Context, requesterId int, addresseeId int) (bool, error)
}
