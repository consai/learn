package handler

import (
	JwtHandler "blog/internal/token"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	userGroup(r.Group("/users"))

	api := r.Group("/api/")
	api.Use(JwtHandler.JWTAuth())

	ApiGroup(api)
}
