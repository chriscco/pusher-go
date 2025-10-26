package domain

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ModelCallRequest struct {
	Model   string `json:"model"`
	ApiKey  string `json:"api_key"`
	Content string `json:"content"`
}

type ModelRequest struct {
	Model    string    `json:"model"`
	ApiKey   string    `json:"api_key"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Message Message `json:"message"`
}

type ModelResponse struct {
	Answer  string         `json:"answer"`
	Choices []Choice       `json:"choices"`
	Error   map[string]any `json:"error"`
}
