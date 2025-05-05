package entities

import "github.com/google/uuid"

type (
	InviteToken struct {
		ID      int       `db:"id"`
		Content uuid.UUID `db:"content"`
		Used    bool      `db:"used"`
		Role    string    `db:"role"`
	}
)
