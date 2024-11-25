package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Ready API
// @Description  Returns a simple message indicating that the service is ready.
// @Tags         readyCheck
// @Accept       json
// @Produce      json
// @Success      202  {string}  string  "Hello Pullu"
// @Router       /ready [get]
func Ready(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "Hello! I am Ready 😄")
}
