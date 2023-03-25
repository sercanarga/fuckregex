package openai_model

type Response struct {
	ID    string `json:"id"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		TotalToken int64 `json:"total_tokens"`
	} `json:"usage"`
	ResponseTime int64 `json:"created"`
}
