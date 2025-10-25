package domain

type SaveRequest struct {
	FileName string `json:"file_name"`
	Content  string `json:"content"`
}

type ReadRequest struct {
	FileName string `json:"file_name"`
}
