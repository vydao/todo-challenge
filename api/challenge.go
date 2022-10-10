package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/vydao/todo-challenge/db/sqlc"
	"github.com/vydao/todo-challenge/token"
)

const dateLayout = "2006-01-02T15:04:05"

type createChallengeRequest struct {
	StartDate   string `json:"start_date" binding:"required"`
	Name        string `json:"name" binding:"required,min=1"`
	Description string `json:"description"`
}

type challengeResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	Description string    `json:"description"`
	UserID      int64     `json:"user_id"`
}

func (sv *Server) CreateChallengeHandler(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var req createChallengeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	startDate, err := time.Parse(dateLayout, req.StartDate)
	if err != nil || startDate.IsZero() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	chalParams := db.CreateChallengeParams{
		StartDate:   startDate,
		Name:        req.Name,
		Description: req.Description,
		UserID:      authPayload.UserID,
	}
	chal, err := sv.store.CreateChallenge(ctx, chalParams)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := challengeResponse{
		ID:          chal.ID,
		Name:        chal.Name,
		Description: chal.Description,
		UserID:      chal.UserID,
		StartDate:   chal.StartDate,
	}
	ctx.JSON(http.StatusOK, resp)
}
