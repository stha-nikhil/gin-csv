package api

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
)

type FileUpload struct {
	CSVFile *multipart.FileHeader `form:"file"`
}

type User struct {
	Name    string
	Age     string
	Address string
}

func GetCSV(c *gin.Context) {
	var csvfile FileUpload
	if err := c.ShouldBind(&csvfile); err != nil {
		log.Fatalln("Error in JSON binding: ", err)

	}

	if csvfile.CSVFile == nil {
		log.Fatalln("File is missing")
	}

	file, err := csvfile.CSVFile.Open()
	if err != nil {
		log.Fatalln("Error in opening file")
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln("Error in reading file")
	}

	var users []User
	for _, record := range records {
		user := User{
			Name:    record[0],
			Age:     record[1],
			Address: record[2],
		}
		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}
