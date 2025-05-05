package models

import "github.com/golang-jwt/jwt"

type (
	ClaimsJwt struct {
		UserID int
		Login  string
		Role   string
		jwt.StandardClaims
	}
	PostInviteTokenRequest struct {
		Role string `json:"role"`
	}
	PostInviteTokenResponse struct {
		IviteToken string `json:"invite_token"`
	}
	AdminSignInRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	AdminSignInResponse struct {
		Token string `json:"token"`
	}
)
