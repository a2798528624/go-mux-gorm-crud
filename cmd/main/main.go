package main

import (
	"go-mux-gorm-crud/internal/model"
	"go-mux-gorm-crud/internal/route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initEnv() {
	err := godotenv.Load("deploy/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnv()
	model.InitDB()
	r := mux.NewRouter()

	// 注册路由
	route.RegisterPlaceRoute(r)
	http.Handle("/", r)

	log.Println("Listening...", "http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
