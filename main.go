package main

import (
	"mock-ses/config"
	"mock-ses/handlers"
	"mock-ses/stats"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	mockSES := &config.MockSESConfig{
		InSandboxMode:   true,
		DailyQuota:      200,
		SendRate:        1,
		VerifiedEmails:  make(map[string]bool),
		VerifiedDomains: make(map[string]bool),
		Statistics: &stats.Statistics{
			DailyCount: make(map[string]int),
			Recipients: make(map[string]int),
			Templates:  make(map[string]int),
		},
	}

	r.POST("/v2/email/send", handlers.SendEMailHandler(mockSES))
	r.GET("/stats", handlers.GetStatsHandler(mockSES))

	r.Run(":8080")
}
