package tracks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRecommendation(c *gin.Context) {
	ctx := c.Request.Context()
	trackID := c.Query("trackID")
	LimitStr := c.Query("limit")
	limit, err := strconv.Atoi(LimitStr)
	if err != nil {
		limit = 10
	}
	userID := c.GetUint("userID")
	response, err := h.service.GetRecommendation(ctx, userID, limit, trackID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, response)
}
