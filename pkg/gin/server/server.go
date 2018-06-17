package server

import (
	"github.com/syunkitada/go-sample/pkg/gin/server/handler"
	"log"
	"net/http"
	"time"
)

func Main() {
	app_handler := handler.GetHandler()

	s := &http.Server{
		Addr:           ":8000",
		Handler:        app_handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
