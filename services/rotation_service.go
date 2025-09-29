package services

import (
	"context"
	"fmt"
	"slices"

	"github.com/aloussase/squad-rotation-bot/entities"
	"github.com/jackc/pgx/v5"
)

type RotationService interface {
	/// Choose the next member in rotation. If the input list is empty, an error is returned.
	ChooseNextInRotation(members []entities.SquadMember) (entities.SquadMember, error)
}

type rotationServiceImpl struct {
	conn *pgx.Conn
}

func CreateRotationService(conn *pgx.Conn) RotationService {
	return &rotationServiceImpl{
		conn: conn,
	}
}

func (rs *rotationServiceImpl) ChooseNextInRotation(members []entities.SquadMember) (entities.SquadMember, error) {
	if len(members) == 0 {
		return entities.SquadMember{}, fmt.Errorf("list of members if empty")
	}

	var next int

	err := rs.conn.QueryRow(context.Background(), "select nextval('squad_rotation')").Scan(&next)
	if err != nil {
		return entities.SquadMember{}, err
	}

	// Sort the members to ensure a deterministic pick.
	sortedMembers := slices.SortedFunc(slices.Values(members), func(a, b entities.SquadMember) int {
		return a.ID - b.ID
	})

	chosen := sortedMembers[next%len(sortedMembers)]

	return chosen, nil
}
