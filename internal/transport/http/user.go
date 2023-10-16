package http

import (
	"database/sql"
	"net/http"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore"
	"github.com/gin-gonic/gin"
)

func newUserResponse(user domain.User) domain.UserResponse {
	return domain.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// @Summary     Signup
// @Description Signup new user
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param 		EnterDetails	body 	domain.SignupRequest true "Signup"
// @Success     200 {object} domain.SignupResponse
// @Failure     500 {object} ErrorResponse
// @Router      /users/signup [post]
func (h *Handler) signUp(ctx *gin.Context) {
	var req domain.SignupRequest

	err := ctx.ShouldBindJSON(&req)

	if handleBindErr(ctx, err) {
		return
	}

	user, err := h.service.User.CreateUser(ctx, req.Email, req.Password, req.Name)
	if err != nil {
		if pgstore.ErrorCode(err) == pgstore.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, err := h.tokenMaker.CreateToken(user.ID, user.Email, h.cfg.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := domain.SignupResponse{
		AccessToken: token,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary     Login
// @Description Login user
// @ID          login
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param 		EnterDetails	body 	domain.LoginRequest true "Login"
// @Success     200 {object} domain.LoginResponse
// @Failure     500 {object} ErrorResponse
// @Router      /users/login [post]
func (h *Handler) login(ctx *gin.Context) {
	var req domain.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := h.service.User.GetUserByEmail(ctx, req.Email)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, err := h.tokenMaker.CreateToken(user.ID, user.Email, h.cfg.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := domain.LoginResponse{
		AccessToken: token,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, res)
}
