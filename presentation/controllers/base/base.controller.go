package base_controller

import "github.com/gin-gonic/gin"

type BaseController interface {
	RegisterRoutes(r *gin.Engine)
}
