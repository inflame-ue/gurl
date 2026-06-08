package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/inflame-ue/gurl/internal/api"
	"github.com/inflame-ue/gurl/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env file: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("could not convert port value to int: %v", err)
	}
	dbDriver := os.Getenv("DB_DRIVER")
	dbDataSource := os.Getenv("DB_DATASOURCE")

	db, err := database.NewDatabase(dbDriver, dbDataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()
	srv := api.NewServer(db)

	serv := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      srv.Mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("server started on port %d", port)
	err = serv.ListenAndServe()
	log.Fatal(err)
}
