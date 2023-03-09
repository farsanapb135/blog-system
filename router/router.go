package router

import (
	"app/blog-system/config"
	"app/blog-system/handler"
	"app/blog-system/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func Routes() http.Handler {

	standardMiddleware := alice.New(middleware.SecureHeaders)

	//definitions of the available API
	router := mux.NewRouter()
	router.HandleFunc("/articles", handler.ArticleManagement).Methods("GET")
	router.HandleFunc("/articles", handler.ArticleManagement).Methods("POST")
	router.HandleFunc("/articles/{article_id}", handler.GetArticlesByID).Methods("GET")

	// viper.SetDefault("listen.address", "localhost")
	// viper.SetDefault("listen.port", "8080")
	listenHost := config.Data.ListenAddrs
	listenPort := config.Data.Port
	socket := listenHost + ":" + listenPort

	fmt.Println("API Started - Listening on: ", socket)

	err := http.ListenAndServe(socket, router)
	if err != nil {
		panic(err)
	}
	return standardMiddleware.Then(router)

}
