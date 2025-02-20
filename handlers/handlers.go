package handlers

import (
	"mock-ses/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatsHandler(mockSES *config.MockSESConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		stats := mockSES.GetStats()
		c.JSON(http.StatusOK, stats)
	}
}

func SendEMailHandler(mockSES *config.MockSESConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		mockSES.SendEmail(c)
	}
}
