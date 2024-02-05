package metrics

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler returns an health `http.Handler`.
func HealthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Timestamp time.Time `json:"timestamp"`
			Status    string    `json:"status"`
		}{
			Timestamp: time.Now(),
			Status:    "ok",
		})
	}
}
