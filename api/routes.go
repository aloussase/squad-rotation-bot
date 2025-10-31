package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aloussase/squad-rotation-bot/services"
)

func ListMembers(
	memberService services.MemberService,
	w http.ResponseWriter,
	r *http.Request,
) {
	members, err := memberService.ListMembers()
	if err != nil {
		log.Printf("Failed to list members: %v", err)
		w.WriteHeader(400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func CreateMember(
	memberService services.MemberService,
	w http.ResponseWriter,
	r *http.Request,
) {
	var body struct {
		FullName  string
		AvatarUrl string
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = memberService.CreateMember(body.FullName, body.AvatarUrl)
	if err != nil {
		log.Printf("Failed to create member: %v", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
}

func TriggerBot(
	memberService services.MemberService,
	rotationService services.RotationService,
	messagingService services.MessagingService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		log.Printf("Succesfully sent rotation update: %s", chosenOne.FullName)
	}
}
