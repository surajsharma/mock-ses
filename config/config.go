package config

import (
	"mock-ses/email"
	"mock-ses/stats"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MockSESConfig struct {
	InSandboxMode   bool
	DailyQuota      int
	SendRate        float64
	VerifiedEmails  map[string]bool
	VerifiedDomains map[string]bool
	Statistics      *stats.Statistics
}

func (m *MockSESConfig) isEmailVerified(input string) bool {
	return false
}

func (m *MockSESConfig) checkQuota() bool {
	return false
}

func (m *MockSESConfig) checkRateLimit() bool {
	return false
}

func (m *MockSESConfig) GetStats() map[string]interface{} {
	return m.GetStats()
}

func (m *MockSESConfig) SendEmail(c *gin.Context) {
	var input email.SendEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error":   "ValidationError",
			"message": err.Error(),
		})
		return
	}

	// Verify sender
	if !m.isEmailVerified(input.Source) {
		c.JSON(400, gin.H{
			"error":   "EmailNotVerified",
			"message": "Email address is not verified",
		})
		return
	}

	// Check quota
	if !m.checkQuota() {
		c.JSON(400, gin.H{
			"error":   "DailyQuotaExceeded",
			"message": "Daily sending quota exceeded",
		})
		return
	}

	// Check rate limiting
	if !m.checkRateLimit() {
		c.JSON(400, gin.H{
			"error":   "MaxSendingRateExceeded",
			"message": "Maximum sending rate exceeded",
		})
		return
	}

	// Record statistics
	m.Statistics.RecordSend(input)

	c.JSON(200, gin.H{
		"MessageId": generateMessageId(),
		"RequestId": generateRequestId(),
	})
}

func generateMessageId() string {
	return uuid.New().String()
}

func generateRequestId() string {
	return uuid.New().String()
}
