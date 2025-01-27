package tracks

import (
	"api-music/internal/models/trackacktivities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertTrackActivities(c *gin.Context) {
	ctx := c.Request.Context()

	var req trackacktivities.TrackActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	UserID := c.GetUint("userID")
	err := h.service.UpsertTrackActivities(ctx, UserID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)

}
