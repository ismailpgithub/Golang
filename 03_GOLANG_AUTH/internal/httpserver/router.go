package httpserver

import (
	// "encoding/json"
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)
	userSvc := user.NewService(userRepo, a.Config.JWTSecret)
	userHandler := user.NewHandler(userSvc)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	api := r.Group("/api")

	api.Use(middleware.AuthRequired(a.Config.JWTSecret))

	api.GET("/files", func(c *gin.Context) {
		usedId, _ := middleware.GetUserID(c)
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"userId": usedId,
			"files":  []any{},
		})
	})

	api.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":       true,
			"products": []any{},
		})
	})

	admin := api.Group("/admin")

	admin.Use(middleware.RequireAdmin())

	admin.GET("/restricted", func(ctx *gin.Context) {
		role, _ := middleware.GetRole(ctx)
		ctx.JSON(http.StatusOK, gin.H{
			"ok":   "true",
			"role": role,
		})
	})

	return r
}
