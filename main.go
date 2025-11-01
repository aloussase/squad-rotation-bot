package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aloussase/squad-rotation-bot/api"
	"github.com/aloussase/squad-rotation-bot/config"
	"github.com/aloussase/squad-rotation-bot/services"
	"github.com/jackc/pgx/v5"
)

func connectDB(dbUrl string) *pgx.Conn {
	attempts := 0

	var conn *pgx.Conn
	var err error

	for attempts < 3 {
		conn, err = pgx.Connect(context.Background(), dbUrl)
		if err == nil {
			break
		}
		if attempts == 2 {
			log.Fatalf("There was an error while trying to connect to the database: %s", err)
		}
		log.Printf("Failed to connect to database, retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
	}

	return conn
}

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("There was an error while reading the config: %s", err)
	}

	conn := connectDB(config.DatabaseUrl)
	defer conn.Close(context.Background())

	memberService := services.Create(conn)
	rotationService := services.CreateRotationService(conn)
	messagingService := services.CreateMessagingService(config)

	http.HandleFunc("/api/v1/rotation/trigger", api.TriggerBot(memberService, rotationService, messagingService))
	http.HandleFunc("/api/v1/rotation/members", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.ListMembers(memberService, w, r)
		case http.MethodPost:
			api.CreateMember(memberService, w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
