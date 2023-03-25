package api_model

type Generate struct {
	Desc string  `json:"desc"`
	Type int     `json:"type"`
	Lang *string `json:"lang"`
}

type Get struct {
	ID string `json:"id"`
}
