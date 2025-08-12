package interfaces

import (
	"Architecture/internal/interfaces/http"
	"Architecture/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, userService usecase.UserService) *gin.Engine {
	userHandler := http.NewUserHandler(userService)

	// User endpointlari
	r.GET("/api/users", userHandler.Get)
	r.GET("/api/users/:id", userHandler.GetByID)
	r.POST("/api//users", userHandler.Create)
	r.PUT("/api/users/:id", userHandler.Update)
	r.DELETE("api/users/:id", userHandler.Delete)

	return r
}
