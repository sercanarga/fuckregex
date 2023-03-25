package api_model

type APIResponse struct {
	ResponseID    string
	InputText     string
	ResponseText  string
	ResponseToken int64
	ResponseTime  int64
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
	ResponseTime int64
}
