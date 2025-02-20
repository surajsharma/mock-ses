package email

// SendEmailInput represents the input structure for sending an email
// swagger:model
type SendEmailInput struct {
	// Source email address
	Source string `json:"Source" binding:"required" example:"sender@example.com"`
	// Destination contains the recipient email addresses
	Destination struct {
		// ToAddresses is a list of primary recipients
		ToAddresses []string `json:"ToAddresses" example:"['recipient1@example.com','recipient2@example.com']"`
		// CcAddresses is a list of CC recipients
		CcAddresses []string `json:"CcAddresses" example:"['cc1@example.com','cc2@example.com']"`
		// BccAddresses is a list of BCC recipients
		BccAddresses []string `json:"BccAddresses" example:"['bcc1@example.com','bcc2@example.com']"`
	} `json:"Destination"`
	// Message contains the email content
	Message struct {
		// Subject of the email
		Subject struct {
			// Data is the subject text
			Data string `json:"Data" example:"Email Subject"`
			// Charset for the subject
			Charset string `json:"Charset" example:"UTF-8"`
		} `json:"Subject"`
		// Body of the email
		Body struct {
			// Text version of the email
			Text struct {
				// Data is the plain text content
				Data string `json:"Data" example:"Plain text email content"`
				// Charset for the text content
				Charset string `json:"Charset" example:"UTF-8"`
			} `json:"Text"`
			// Html version of the email
			Html struct {
				// Data is the HTML content
				Data string `json:"Data" example:"<html><body>HTML email content</body></html>"`
				// Charset for the HTML content
				Charset string `json:"Charset" example:"UTF-8"`
			} `json:"Html"`
		} `json:"Body"`
	} `json:"Message"`
}
