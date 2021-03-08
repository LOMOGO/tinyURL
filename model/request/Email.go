package request

type EmailInfo struct {
	Recipient string `json:"recipient" binding:"email"`
	URL       string `json:"url" binding:"url"`
}
