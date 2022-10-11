package service

import (
	"context"
	"database/sql"
	"github.com/quangpham789/golang-assessment/repository"
	"github.com/quangpham789/golang-assessment/repository/relation"
	"github.com/quangpham789/golang-assessment/repository/user"
)

type RelationsService struct {
	relationsRepository repository.RelationsRepo
	userRepository      repository.UserRepo
}

func (serv RelationsService) CreateRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error) {
	//get requesterId from email request
	requesterId, err := serv.userRepository.GetUserByEmail(ctx, input.RequesterEmail)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	//get addresseeId from email request
	addresseeId, err := serv.userRepository.GetUserByEmail(ctx, input.AddresseeEmail)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	// insert relation two email
	result, err := serv.relationsRepository.CreateRelation(ctx, requesterId.ID, addresseeId.ID)
	if err != nil {
		return CreateRelationsResponse{Success: false}, err
	}

	return CreateRelationsResponse{Success: result}, nil
}

type CreateRelationsInput struct {
	RequesterEmail string
	AddresseeEmail string
}

type CreateRelationsResponse struct {
	Success bool `json:"success"`
}

type RelationServ interface {
	CreateRelation(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error)
}

func NewRelationService(db *sql.DB) RelationServ {
	return RelationsService{
		relationsRepository: relation.NewRelationsRepository(db),
		userRepository:      user.NewUserRepository(db),
	}
}
