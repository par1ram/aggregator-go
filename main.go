package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in enviroment ")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1 := chi.NewRouter()
	routerV1.Get("/healthz", handlerReadiness)
	routerV1.Get("/error", handlerError)
	router.Mount("/v1", routerV1)

	server := http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Server started on PORT:", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
