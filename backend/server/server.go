package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ArthurPaivaT/Porletu/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Start(resCh chan error) {
	fmt.Println("Starting Server...")

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	router := mux.NewRouter()

	router.Handle("/register", corsWrapper.Handler(http.HandlerFunc(handlers.CreateUser))).Methods("POST")
	router.Handle("/deleteUser", corsWrapper.Handler(http.HandlerFunc(handlers.DeleteUser))).Methods("POST")

	fmt.Println("Listening on Port :1212")
	err := http.ListenAndServe(":1212", router)
	if err != nil {
		fmt.Println("Could not start server:", err)
	}

	resCh <- errors.New("Erro ao iniciar servidor")
}
