package services

import (
	"context"

	"github.com/aloussase/squad-rotation-bot/entities"
	"github.com/jackc/pgx/v5"
)

type MemberService interface {
	// ListMembers / List all members of the squad.
	ListMembers() ([]entities.SquadMember, error)

	// Create a new member.
	CreateMember(name, avatarUrl string) error
}

type memberServiceImpl struct {
	conn *pgx.Conn
}

// Create / Create a new instance of MemberService.
func Create(conn *pgx.Conn) MemberService {
	return &memberServiceImpl{
		conn: conn,
	}
}

func (ms *memberServiceImpl) CreateMember(name, avatarUrl string) error {
	query := "insert into squad_members (full_name, avatar_url) values ($1, $2)"
	_, err := ms.conn.Exec(context.Background(), query, name, avatarUrl)
	return err
}

func (ms *memberServiceImpl) ListMembers() ([]entities.SquadMember, error) {
	query := "select (id, full_name, avatar_url) from squad_members"
	rows, err := ms.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	members, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (entities.SquadMember, error) {
		var member entities.SquadMember
		err := row.Scan(&member)
		return member, err
	})

	if err != nil {
		return nil, err
	}

	return members, nil
}
