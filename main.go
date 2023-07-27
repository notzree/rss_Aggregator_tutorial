package main

import (
	//"fmt"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}
	dsn := fmt.Sprintf("%s&parseTime=True", os.Getenv("DSN"))
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if dbErr != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", dbErr)
	}
	log.Println("Successfully connected to PlanetScale!")

	migrateErr := db.AutoMigrate(&Users{}, &Feeds{})

	if migrateErr != nil {
		log.Fatalf("failed to migrate: %v", migrateErr)
	}
	handler := newHandler(db)

	//main router, think of a router like resource in aws api gateway
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//creating a sub router
	v1Router := chi.NewRouter()
	//Attatching a function that will respond to the /healthz path (used to check if server is running)
	//using .Get (over .HandleFunc) scopes the route to only GET requests
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	v1Router.Post("/createUser", handler.handlerCreateUser)
	v1Router.Post("/createFeed", handler.handleCreateFeed)
	v1Router.Post("/getFeed", handler.handleReadFeed)
	v1Router.Post("/getUser", handler.handleGetUser)
	//nesting the router under the /v1 path
	//This means that the final route for the /ready path is: /v1/healthz
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server started on port %s", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}
