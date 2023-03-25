package db_model

type Responses struct {
	ID            string
	InputText     string
	ResponseText  string
	ResponseToken int64
	CreatedDate   int64
	IsReported    bool
}
