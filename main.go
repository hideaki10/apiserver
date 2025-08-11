package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hideaki10/apiserver/handlers"
)

func main() {

	r := chi.NewRouter()

	r.Get("/hello", handlers.HelloHandler)
	r.Get("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ArticlelistHandler)

	r.Get("/article/l", handlers.ArticleDetailHandler)
	r.Get("/article/nice", handlers.PostNiceHandler)
	r.Get("/article/nice", handlers.PostNiceHandler)
	r.Post("/comment", handlers.CommentHandler)

	// http.HandleFunc("/hello", handlers.HelloHandler)
	// http.HandleFunc("/article", handlers.PostArticleHandler)
	// http.HandleFunc("/article/list", handlers.ArticlelistHandler)
	// http.HandleFunc("/article/l", handlers.ArticleDetailHandler)
	// http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// http.HandleFunc("/comment", handlers.CommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
