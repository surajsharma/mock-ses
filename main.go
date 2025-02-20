package main

import (
	"mock-ses/config"
	"mock-ses/handlers"
	"mock-ses/stats"

	docs "mock-ses/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/v2"

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/v2/email/SendEmail", handlers.SendEMailHandler(mockSES))
	r.GET("/v2/email/GetSendStatistics", handlers.GetStatsHandler(mockSES))

	r.Run(":8080")
}
