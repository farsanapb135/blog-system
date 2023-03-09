package handler

import (
	"app/blog-system/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ArticleManagement(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		var articles model.Article

		values, err := model.GetArticleDetailsFromDB()
		if err != nil {
			articles.Status = 400
			articles.Message = err.Error()
			articles.Data = nil

		} else {
			articles.Status = 200
			articles.Message = "Success"
			articles.Data = values.Data
		}

		sliceForArticle, _ := json.Marshal(articles)
		w.Write([]byte(sliceForArticle))

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("curl body error", err)
		}

		clientData := model.ArticleData{}
		responseData := model.ArticleResult{}
		err = json.Unmarshal(body, &clientData)
		if err != nil {
			responseData.Status = 400
			responseData.Message = "Bad Request"
			responseData.Data.ID = "nil"
			sliceforApiResponse, _ := json.Marshal(responseData)
			w.Write([]byte(sliceforApiResponse))
		} else {

			if clientData.Author != "" && clientData.Content != "" && clientData.Title != "" {
				errData, id := model.AddArticleToDb(clientData)
				if errData != nil {
					responseData.Status = 400
					responseData.Message = errData.Error()
					responseData.Data.ID = "nil"
				} else {
					responseData.Status = 201
					responseData.Message = "Success"
					responseData.Data.ID = id

				}
			} else {
				responseData.Status = 400
				responseData.Message = "Bad Request"
				responseData.Data.ID = "nil"
			}
			sliceForArticle, _ := json.Marshal(responseData)
			w.Write([]byte(sliceForArticle))

		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func GetArticlesByID(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var articles model.Article
		params := mux.Vars(r)
		id := params["article_id"]

		if id == "" {
			articles.Status = 400
			articles.Message = "Bad Request"
			articles.Data = nil
		}

		values, err := model.GetArticleDetailsByID(id)
		if err != nil {
			articles.Status = 400
			articles.Message = err.Error()
			articles.Data = nil

		} else {
			articles.Status = 200
			articles.Message = "Success"
			articles.Data = values.Data
		}

		sliceFOrArticle, _ := json.Marshal(articles)
		w.Write([]byte(sliceFOrArticle))
	}

}
