package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ybencab/url-shortener/api"
	"github.com/ybencab/url-shortener/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	url := os.Getenv("DB_URL")
	store, err := storage.NewPostgresStorage(url)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	listenAddr := os.Getenv("PORT")
	log.Println("server running on port ", listenAddr)
	
	server := api.NewServer(listenAddr, store)
	if err := server.Run(); err != nil {
		log.Fatal("error: ", err.Error())
	}
}
