package api

import (
	"database/sql"
	"net/http"
	"time"

	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/arya2004/Xyfin/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Passowrd string `json:"password" binding:"required,min=8"`
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userDto struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Passowrd string `json:"password" binding:"required,min=8"`

}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	User userDto `json:"user"`

}

func newUserReponse(user database.User) userDto {
	return userDto{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		PasswordChangedAt : user.PasswordChangedAt,
		CreatedAt : user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {

	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Passowrd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := database.CreateUserParams{
		FullName: req.FullName,
		Email: req.Email,
		Username: req.Username,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	dto := newUserReponse(user)

	ctx.JSON(http.StatusOK, dto)

}


func (server *Server) loginUser(ctx *gin.Context){
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	err = util.CheckPassword(req.Passowrd, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenCreator.CreateToken(
		user.Username,
		server.configuration.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User: newUserReponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)

}