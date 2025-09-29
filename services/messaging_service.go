package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aloussase/squad-rotation-bot/config"
	"github.com/aloussase/squad-rotation-bot/entities"
)

type MessagingService interface {
	SendRotationNotification(member entities.SquadMember) error
}

type messagingServiceImpl struct {
	config *config.Config
}

func CreateMessagingService(config *config.Config) MessagingService {
	return &messagingServiceImpl{config}
}

func (m *messagingServiceImpl) SendRotationNotification(member entities.SquadMember) error {
	payload := map[string]any{
		"cards": map[string]any{
			"header": map[string]any{
				"title":    "Today's Presenter 🎤",
				"imageUrl": member.AvatarUrl,
			},
			"sections": []any{
				map[string]any{
					"widgets": []any{
						map[string]any{
							"textParagraph": map[string]any{
								"text": member.FullName,
							},
						},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(m.config.WebHookUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
