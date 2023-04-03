package api

import (
	"mime/multipart"
)

type FileUpload struct {
	CSVFile *multipart.FileHeader `form:"file"`
}

type User struct {
	Name    string
	Age     string
	Address string
}
