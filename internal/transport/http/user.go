package http

import (
	"database/sql"
	"net/http"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/token"
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

func (server *Handler) listUserUrls(ctx *gin.Context) {
	var req domain.ListUserUrlsRequest

	err := ctx.ShouldBindJSON(&req)

	if handleBindErr(ctx, err) {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	urls, err := server.service.User.ListUserUrls(ctx, authPayload.UserID, req.Limit, req.Offset)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := domain.ListUserUrlsResponse{
		UserID: authPayload.UserID,
		Urls:   urls,
	}

	ctx.JSON(http.StatusOK, res)
}
