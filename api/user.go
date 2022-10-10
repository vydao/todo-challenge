package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/vydao/todo-challenge/db/sqlc"
)

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type UserResponse struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}

func (sv *Server) GetUserHandler(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := sv.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Score:    user.Score,
	}})
}

func (sv *Server) CreateUserHandler(ctx *gin.Context) {
	userReq := db.CreateUserParams{}
	if err := ctx.BindJSON(&userReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	_, err := sv.store.CreateUser(ctx, userReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}

func (sv *Server) LoginUserHandler(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := sv.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// TODO: should use hashed password
	if req.Password != user.Password {
		err = errors.New("the username or password is incorrect")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	// TODO: move duration to config
	token, err := sv.tokenMaker.CreateToken(user, time.Minute*30)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, loginUserResponse{token})
}
