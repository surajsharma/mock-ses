package config

import (
	"mock-ses/email"
	"mock-ses/stats"
	"strings"

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
	//any verification logic

	allowedDomains := []string{"gmail.com", "hotmail.com", "yahoo.com"}

	parts := strings.Split(input, "@")

	if len(parts) != 2 {
		return false
	}

	domain := strings.ToLower(parts[1])

	for _, allowedDomain := range allowedDomains {
		if domain == allowedDomain {
			return true
		}
	}

	return false
}

func (m *MockSESConfig) checkQuota() bool {
	return m.Statistics.EmailsSent <= m.DailyQuota
}

func (m *MockSESConfig) checkRateLimit() bool {
	return m.Statistics.EmailsSent <= int(m.SendRate)
}

func (m *MockSESConfig) GetStats() map[string]interface{} {
	return m.Statistics.GetStats()
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
