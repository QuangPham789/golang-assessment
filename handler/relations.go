package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/quangpham789/golang-assessment/service"
	"github.com/quangpham789/golang-assessment/utils"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

type RelationsHandler struct {
	relationsService service.RelationServ
}

func NewRelationsHandler(db *sql.DB) RelationsHandler {
	return RelationsHandler{
		relationsService: service.NewRelationService(db),
	}
}

type RelationRequest struct {
	Friends []string `json:"friends"`
}

func (relations RelationsHandler) CreateRelation(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	relationReq := RelationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&relationReq); err != nil {
		JsonResponseError(w, err)
		return
	}

	// Validate body user request
	input, err := validateFriendshipInput(relationReq)
	if err != nil {
		JsonResponseError(w, err)
		return
	}

	// check user block

	// Call service
	result, err := relations.relationsService.CreateRelation(r.Context(), input)
	if err != nil {
		log.Println("CreateRelation error", err)
		utils.JsonResponse(w, http.StatusForbidden, result)
	}
	utils.JsonResponse(w, http.StatusCreated, result)
}

func validateFriendshipInput(relationReq RelationRequest) (service.CreateRelationsInput, error) {
	//check of slice of friend is empty
	if len(relationReq.Friends) < 2 {
		return service.CreateRelationsInput{}, errDataIsEmpty
	}
	requesterEmail := strings.TrimSpace(relationReq.Friends[0])
	if requesterEmail == "" {
		return service.CreateRelationsInput{}, errEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(requesterEmail); err != nil {
		return service.CreateRelationsInput{}, errInvalidEmail
	}

	addresseeEmail := strings.TrimSpace(relationReq.Friends[1])
	if addresseeEmail == "" {
		return service.CreateRelationsInput{}, errEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(addresseeEmail); err != nil {
		return service.CreateRelationsInput{}, errInvalidEmail
	}

	return service.CreateRelationsInput{
		RequesterEmail: requesterEmail,
		AddresseeEmail: addresseeEmail,
	}, nil
}
