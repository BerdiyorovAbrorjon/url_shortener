package http

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) createUrl(ctx *gin.Context) {
	var req domain.CreateUrlRequest

	err := ctx.ShouldBindJSON(&req)
	if handleBindErr(ctx, err) {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	url, err := h.service.Url.CreateUrl(ctx, req.OrgUrl, authPayload.UserID)

	if err != nil {
		switch pgstore.ErrorCode(err) {
		case pgstore.UniqueViolation, pgstore.ForeignKeyViolation:
			{
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, url)
}

func (h *Handler) getUrlById(ctx *gin.Context) {
	var req domain.GetUrlByIdRequest

	err := ctx.ShouldBindUri(&req)
	if handleBindErr(ctx, err) {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	url, err := h.service.Url.GetUrlById(ctx, req.ID)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if url.UserID != authPayload.UserID {
		err = errors.New("another user error. user can get only own url")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, url)
}

func (h *Handler) updateOrgUrl(ctx *gin.Context) {
	var req domain.UpdateOrgUrlRequest

	err := ctx.ShouldBindJSON(&req)
	if handleBindErr(ctx, err) {
		return
	}

	url, err := h.service.Url.GetUrlById(ctx, req.ID)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(token.Payload)

	if url.UserID != authPayload.UserID {
		err = errors.New("another user error. user can update only own url")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	updUrl, err := h.service.Url.UpdateOrgUrl(ctx, req.ID, req.NewOrgUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updUrl)
}

func (h *Handler) deleteUrl(ctx *gin.Context) {
	var req domain.DeleteUrlRequest

	err := ctx.ShouldBindUri(&req)
	if handleBindErr(ctx, err) {
		return
	}

	url, err := h.service.Url.GetUrlById(ctx, req.ID)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(token.Payload)

	if url.UserID != authPayload.UserID {
		err = errors.New("another user error. user can delete only own url")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = h.service.Url.DeleteUrl(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) rederect(ctx *gin.Context) {
	var req domain.RederectRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	log.Info().Msg("Rederct input:" + req.ShortUrl)

	orgUrl, err := h.service.Url.GetOrgUrlByShort(ctx, req.ShortUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, orgUrl)
}
