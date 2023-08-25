package utils

type ErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
type SuccessResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
}

type Data struct {
	Filename string `json:"filename"`
	FileExt  string `json:"fileext"`
	Filesize int64  `json:"filesize"`
	FileType string `json:"filetype"`
}
