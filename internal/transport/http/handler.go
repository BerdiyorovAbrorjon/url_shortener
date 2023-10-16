package http

import (
	"net/http"

	"github.com/BerdiyorovAbrorjon/url-shortener/config"
	_ "github.com/BerdiyorovAbrorjon/url-shortener/docs"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/service"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/token"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	cfg        config.Config
	service    *service.Service
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewHandler(cfg config.Config, service *service.Service, tokenMaker token.Maker) *Handler {
	return &Handler{
		cfg:        cfg,
		service:    service,
		tokenMaker: tokenMaker,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	// Init gin handler
	if h.cfg.AppMode == config.ProdMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.ContextWithFallback = true

	// Options
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Health
	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	h.setupRouter(router)

	return router
}

// Swagger spec:
// @title       Url Shortener API
// @description Shortening a long url and managing access to it
// @version     1.0
// @host        localhost:8080
// @BasePath    /
func (h *Handler) setupRouter(router *gin.Engine) {

	router.POST("/users/signup", h.signUp)
	router.POST("/users/login", h.login)

	authRouter := router.Group("/").Use(authMiddleware(h.tokenMaker))

	authRouter.POST("/urls", authMiddleware(h.tokenMaker), h.createUrl)
	authRouter.GET("/urls/:id", authMiddleware(h.tokenMaker), h.getUrlById)
	authRouter.GET("/urls", authMiddleware(h.tokenMaker), h.listUserUrls)
	authRouter.POST("/urls/update", authMiddleware(h.tokenMaker), h.updateOrgUrl)
	authRouter.DELETE("/urls/:id", authMiddleware(h.tokenMaker), h.deleteUrl)

	router.GET("/:short_url", h.rederect)

	h.router = router
}
