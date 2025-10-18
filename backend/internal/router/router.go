package router

import (
	"backend/internal/article"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	h := article.NewHandler(db)
	r.POST("/articles", h.Create)
	r.GET("/articles", h.List)
	r.GET("/articles/:id", h.Get)
	r.PUT("/articles/:id", h.Update)
	r.DELETE("/articles/:id", h.Delete)

	return r
}
