package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vydao/todo-challenge/db/sqlc"
)

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetUserResponse struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}

func (sv *Server) GetUserHandler(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := sv.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": GetUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Score:    user.Score,
	}})
}

func (sv *Server) CreateUserHandler(ctx *gin.Context) {
	userReq := db.CreateUserParams{}
	if err := ctx.BindJSON(&userReq); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, err := sv.store.CreateUser(ctx, userReq)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
