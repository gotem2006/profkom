package auth

import (
	"context"
	"profkom/internal/entities"
)

type (
	repository interface {
		CheckUserExist(ctx context.Context, login string) (exist bool, err error)
		CheckInviteToken(ctx context.Context, inviteToken string) (role string, err error)
		InsertUser(ctx context.Context, user *entities.User) (err error)
	}
)
