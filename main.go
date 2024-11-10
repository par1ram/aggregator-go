package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/par1ram/aggregator-go/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in enviroment ")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in enviroment ")
	}

	// Connection to database
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cant connect to database", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	// Start scraping
	go startScraping(db, runtime.NumCPU(), time.Minute)

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
	routerV1.Post("/users", apiCfg.handlerCreateUser)
	routerV1.Get("/users", apiCfg.middleWareAuth(apiCfg.handlerGetUser))
	routerV1.Post("/feeds", apiCfg.middleWareAuth(apiCfg.handlerCreateFeed))
	routerV1.Get("/feeds", apiCfg.handlerGetFeeds)
	routerV1.Post("/feed_follows", apiCfg.middleWareAuth(apiCfg.handlerCreateFeedFollow))
	routerV1.Get("/feed_follows", apiCfg.middleWareAuth(apiCfg.handlerGetFeedFollows))
	routerV1.Delete("/feed_follows/{feedFollowId}", apiCfg.middleWareAuth(apiCfg.handlerDeleteFeedFollow))
	routerV1.Get("/posts", apiCfg.middleWareAuth(apiCfg.hanlerGetPostsForUser))
	router.Mount("/v1", routerV1)

	server := http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Server started on PORT:", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
