package handlers

import (
	"mock-ses/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v2
// GetStatsHandler godoc
// @Summary GetSendStatistics Handler
// @Schemes http
// @Description Provides sending statistics for the current AWS Region. The result is a list of data points, representing the last two weeks of sending activity. Each data point in the list contains statistics for a 15-minute period of time. You can execute this operation no more than once per second.
// @Produce json
// @Success 200
// @Router /v2/email/GetSendStatistics [get]
func GetStatsHandler(mockSES *config.MockSESConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		stats := mockSES.Statistics.GetStats()
		c.JSON(http.StatusOK, stats)
	}
}

// @BasePath /v2
// SendEMailHandler godoc
// @Summary SendEmail Handler
// @Schemes http
// @Description Composes an email message and immediately queues it for sending. In order to send email using the SendEmail  operation, your message must meet the following requirements: The message must be sent from a verified email address or domain. If you attempt to send email using a non-verified address or domain, the operation will result in an \"Email address not verified\" error. If your account is still in the Amazon SES sandbox, you may only send to verified addresses or domains, or to email addresses associated with the Amazon SES Mailbox Simulator.
// @Accept json
// @Produce json
// @Param request body SendEmailInput true "Email configuration" SchemaExample({"Source":"sender@example.com","Destination":{"ToAddresses":["recipient1@example.com"],"CcAddresses":["cc1@example.com"],"BccAddresses":["bcc1@example.com"]},"Message":{"Subject":{"Data":"Email Subject","Charset":"UTF-8"},"Body":{"Text":{"Data":"Plain text content","Charset":"UTF-8"},"Html":{"Data":"<html><body>HTML content</body></html>","Charset":"UTF-8"}}}})
// @Success 200
// @Router /v2/email/SendEmail [post]
func SendEMailHandler(mockSES *config.MockSESConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		mockSES.SendEmail(c)
	}
}
