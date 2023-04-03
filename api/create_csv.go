package api

import (
	"bytes"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateCSV(c *gin.Context) {
	records := [][]string{
		{"name", "age", "nationality"},
		{"Nick Shrestha", "22", "Nepali"},
		{"Jon Green", "25", "American"},
	}

	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(records)

	_, err := c.Writer.Write(csvBuffer.Bytes())
	if err != nil {
		log.Fatalln("Error in writing with context: ", err.Error())
	}
}
