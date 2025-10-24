package domain

type EmailRequest struct {
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
	To       []string `json:"to"`
	From     string   `json:"from"`
	Password string   `json:"password"`
}
