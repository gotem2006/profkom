package entities

import (
	"time"

	"gopkg.in/guregu/null.v2"
)

type (
	User struct {
		ID       int         `db:"id"`
		Role     string      `db:"role"`
		Name     null.String `db:"name"`
		Login    string      `db:"login"`
		Password string      `db:"password"`
		CreateAt time.Time   `db:"created_at"`
	}
)
