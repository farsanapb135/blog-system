package main

import (
	"app/blog-system/config"
	"app/blog-system/router"
	"log"
	"net/http"
	"time"
)

func init() {

	config.SetConfiguration()

}

func main() {

	router := router.Routes()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + config.Data.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	// for run https requests
	//log.Fatal(srv.ListenAndServeTLS(config.Data.HTTPSCertFile, config.Data.HTTPSKeyFile))
}
