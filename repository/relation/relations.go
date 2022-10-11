package relation

import (
	"context"
	"database/sql"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/quangpham789/golang-assessment/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewRelationsRepository(db *sql.DB) repository.RelationsRepo {
	return relationsRepository{
		connection: db,
	}
}

type relationsRepository struct {
	connection *sql.DB
}

func (repo relationsRepository) GetRelationByIds(ctx context.Context, requesterId int, addresseeId int) (models.Relation, error) {
	var relationResult models.Relation
	if err := models.Users(models.RelationWhere.RequesterID.EQ(requesterId), models.RelationWhere.AddresseeID.EQ(addresseeId)).Bind(ctx, repo.connection, &relationResult); err != nil {
		return models.Relation{}, err
	}
	return relationResult, nil
}

func (repo relationsRepository) CreateRelation(ctx context.Context, requesterId int, addresseeId int) (bool, error) {
	var relation = models.Relation{}
	relation.RequesterID = requesterId
	relation.AddresseeID = addresseeId
	if err := relation.Insert(ctx, repo.connection, boil.Infer()); err != nil {
		return false, err
	}
	return true, nil
}
