package router

import (
	"backend/internal/post"

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

	h := post.NewHandler(db)

	r.POST("/article/", h.Create)      // POST /article/
	r.GET("/article/:id", h.Get)       // GET /article/1
	r.PUT("/article/:id", h.Update)    // PUT /article/1
	r.DELETE("/article/:id", h.Delete) // DELETE /article/1

	// For pagination - use query parameters instead of path parameters
	r.GET("/article", h.List) // GET /article?limit=10&offset=0

	return r
}
