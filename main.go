package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint hit: homePage")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {

		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			return
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)

	vars := mux.Vars(r)
	id := vars["id"]

	for i, article := range Articles {

		if article.Id == id {
			article.Content = updatedArticle.Content
			article.Desc = updatedArticle.Desc
			article.Title = updatedArticle.Title
			Articles[i] = article
			json.NewEncoder(w).Encode(article)
			return
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "3", Title: "Hello 3", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
