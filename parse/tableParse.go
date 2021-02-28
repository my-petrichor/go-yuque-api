package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/my-Sakura/go-yuque-api/api"
)

type Table struct {
	Data []Fields `json:"data"`
}

type Fields struct {
	Date     string `json:"date"`
	TagColor string `json:"tag_color"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func parseTabel() Table {
	var fields []Fields = make([]Fields, 0)
	token := "YLN7hYz4iKmWSs1MfyLDrNY2IqZaM2ZabOOmpIAX"
	doc := api.GetDocumentInfo(token, "my-Sakura/doc", "hmi4x3")
	trimSpace := strings.TrimSpace(doc.Data.Body)
	splitN := strings.Split(trimSpace, "\n")
	midMultipleSlice := make([][]string, 0)
	for rowNumber, rowSlice := range splitN {
		midSlice := make([]string, 0)
		if rowNumber <= 1 {
			continue
		}

		splitCow := strings.Split(rowSlice, "|")
		for _, value := range splitCow {
			value = strings.TrimSpace(value)
			if value != "" {
				midSlice = append(midSlice, value)
			}
		}
		midMultipleSlice = append(midMultipleSlice, midSlice)
	}

	for _, v := range midMultipleSlice {
		var field Fields = Fields{
			Date:     v[0],
			TagColor: v[1],
			Title:    v[2],
			Content:  v[3],
		}
		fields = append(fields, field)
	}

	table := Table{
		Data: fields,
	}
	return table
}

func tableHandler(w http.ResponseWriter, r *http.Request) {
	table := parseTabel()
	resp, err := json.Marshal(table)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(resp)
}

func main() {
	http.HandleFunc("/", tableHandler)
	err := http.ListenAndServe(":2000", nil)
	if err != nil {
		log.Println(err)
	}
}
