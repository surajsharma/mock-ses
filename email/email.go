package email

type SendEmailInput struct {
	Source      string `json:"Source" binding:"required"`
	Destination struct {
		ToAddresses  []string `json:"ToAddresses"`
		CcAddresses  []string `json:"CcAddresses"`
		BccAddresses []string `json:"BccAddresses"`
	} `json:"Destination"`
	Message struct {
		Subject struct {
			Data    string `json:"Data"`
			Charset string `json:"Charset"`
		} `json:"Subject"`
		Body struct {
			Text struct {
				Data    string `json:"Data"`
				Charset string `json:"Charset"`
			} `json:"Text"`
			Html struct {
				Data    string `json:"Data"`
				Charset string `json:"Charset"`
			} `json:"Html"`
		} `json:"Body"`
	} `json:"Message"`
}
