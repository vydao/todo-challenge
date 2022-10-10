package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/vydao/todo-challenge/db/sqlc"
)

type createTodoRequest struct {
	Name        string  `json:"name" binding:"required"`
	ChallengeID int64   `json:"challenge_id"`
	Period      string  `json:"period" binding:"required,oneof=daily weekly monthly"`
	Point       float64 `json:"point" binding:"required,gte=1"`
}

type createTodoResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (sv *Server) CreateTodoHandler(ctx *gin.Context) {
	challengeID, err := strconv.ParseInt(ctx.Param("challenge_id"), 10, 64)
	if err != nil || challengeID <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req createTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	todo, err := sv.store.CreateTodo(ctx, db.CreateTodoParams{
		Name:        req.Name,
		ChallengeID: challengeID,
		Period:      req.Period,
		Point:       req.Point,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": createTodoResponse{
		ID:   todo.ID,
		Name: todo.Name,
	}})
}

type getTodosByChallengeRequest struct {
	ID int64 `json:"id" uri:"challenge_id"`
}

func (sv *Server) GetTodosByChallengeHandler(ctx *gin.Context) {
	req := getTodosByChallengeRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	todos, err := sv.store.GetTodosByChallenge(ctx, req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": todos})
}
