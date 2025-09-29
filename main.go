package main

import (
	"context"
	"log"

	"github.com/aloussase/squad-rotation-bot/config"
	"github.com/aloussase/squad-rotation-bot/services"
	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("There was an error while reading the config: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), config.DatabaseUrl)
	if err != nil {
		log.Fatalf("There was an error while trying to connect to the database: %s", err)
	}

	defer conn.Close(context.Background())

	memberService := services.Create(conn)
	rotationService := services.CreateRotationService(conn)
	messagingService := services.CreateMessagingService(config)

	members, err := memberService.ListMembers()
	if err != nil {
		log.Fatalf("There was an error while trying to list members: %s", err)
	}

	chosenOne, err := rotationService.ChooseNextInRotation(members)
	if err != nil {
		log.Fatalf("There was an error while trying to choose next in rotation: %s", err)
	}

	if messagingService.SendRotationNotification(chosenOne) != nil {
		log.Fatalf("There was an error while trying to send a rotation: %s", err)
	}
}
