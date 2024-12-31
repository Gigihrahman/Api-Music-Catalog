package memberships

import (
	"api-music/internal/models/memberships"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req memberships.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.service.Login(req)
	if err != nil {
		fmt.Println("test case")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, memberships.LoginResponse{
		AccesToken: accessToken,
	})
}
