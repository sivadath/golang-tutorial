package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port = os.Getenv("PORT")
func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/articles",GetAllArticles).Methods("GET")
	router.HandleFunc("/articles",CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}",GetArticleById).Methods("GET")
	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port,router))
}
