package server

import (
	"db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CreateResponse struct {
	Status uint16
	Message string
	Data []db.Article
}

const responseMsg = `{"status": %d,"message": %s,"data": %s}`

func CreateArticle(writer http.ResponseWriter, req *http.Request) {
		article := db.Article{}
		err := json.NewDecoder(req.Body).Decode(&article)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(writer, fmt.Sprintf(responseMsg, http.StatusBadRequest, "can't read body", "null"), http.StatusBadRequest)
			return
		}
		ID, err := db.NewBlogger().PushArticle(article)
		if err != nil {
			http.Error(writer,fmt.Sprintf(responseMsg, http.StatusInternalServerError, "no data available in db", "null"),http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		_, _ = writer.Write(
			[]byte(
				fmt.Sprintf(responseMsg,http.StatusCreated, "Success",
					fmt.Sprintf(`{"id": %d}`,
						ID))))

}

func GetArticleById(writer http.ResponseWriter, req *http.Request) {
		ID, err := strconv.ParseInt(req.URL.Path[len(`/articles/`):], 10, 64)
		if err != nil {
			log.Printf("Error parsing ID: %v", err)
			http.Error(writer, fmt.Sprintf(responseMsg, http.StatusBadRequest, "can't read body", "null"), http.StatusBadRequest)
			return
		}
		article, err := db.NewBlogger().GetArticle(ID)
		if err != nil {
			log.Printf("Failed to fetch article. %v", err)
			http.Error(writer,fmt.Sprintf(responseMsg, http.StatusInternalServerError, "no data available in db", "null"),http.StatusInternalServerError)
			return
		}
		jsn, err := json.Marshal(article)
		if err != nil {
			log.Printf("Json parsing failed: %v", err)
			http.Error(writer, fmt.Sprintf(responseMsg, http.StatusNotFound, "resource not found", "null"), http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(
			[]byte(
				fmt.Sprintf(responseMsg, http.StatusOK, "Success",
					fmt.Sprintf(`[%s]`, string(jsn)))))
}

func GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	articles, err := db.NewBlogger().GetAllArticles()
	if err != nil {
		log.Printf("Failed to collect articles list: %v",err)
		http.Error(writer,fmt.Sprintf(responseMsg, http.StatusInternalServerError, "no data available in db", "null"),http.StatusInternalServerError)
		return
	}
	jsn, err := json.Marshal(articles)
	if err != nil {
		log.Printf("Json parsing failed: %v", err)
		http.Error(writer, fmt.Sprintf(responseMsg, http.StatusNotFound, "resource not found", "null"), http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(
		[]byte(
			fmt.Sprintf(responseMsg, http.StatusOK, "Success",
				fmt.Sprintf(`%s`, string(jsn)))))
	return
}
