package factories

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/0x000def42/microshards-go-config/controllers"
	"github.com/gorilla/mux"
)

func NewHttpServer(handlers []controllers.ControllerHttp) *http.Server {
	router := mux.NewRouter()

	for _, handler := range handlers {
		handler.Routes(router)
	}
	l := log.New(os.Stdout, "microshards-config ", log.LstdFlags)
	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		fmt.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	return &s
}
