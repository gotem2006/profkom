package models

import (
	"profkom/internal/entities"
)

type (
	SignUpRequest struct {
		Login       string `json:"login"`
		Password    string `json:"password"`
		InviteToken string `json:"invite_token"`
	}
	SignUpResponse struct {
		JwtToken string `json:"token"`
	}
)

func (r *SignUpRequest) ToEntity() entities.User {
	return entities.User{
		Login:    r.Login,
		Password: r.Password,
	}
}
