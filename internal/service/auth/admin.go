package auth

import (
	"context"
	"errors"
	"fmt"
	"profkom/internal/entities"
	"profkom/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AdminSingUp(ctx context.Context, req models.SignUpRequest) (response models.SignUpResponse, err error) {
	var user entities.User

	err = s.txManager.Do(ctx, func(ctx context.Context) error {
		exist, err := s.repo.CheckUserExist(ctx, req.Login)
		if err != nil {
			return err
		}

		if exist {
			return fmt.Errorf("login exist")
		}

		user = req.ToEntity()

		if req.InviteToken == "" {
			return fmt.Errorf("invalid token")
		}

		user.Role, err = s.repo.CheckInviteToken(ctx, req.InviteToken)
		if err != nil {
			return err
		}

		user.Password, err = s.hashPassword(ctx, req.Password)
		if err != nil {
			return err
		}

		err = s.repo.InsertUser(ctx, &user)
		if err != nil {
			return err
		}

		return err
	})
	if err != nil {
		return response, err
	}

	claims := models.ClaimsJwt{
		UserID: user.ID,
		Login:  user.Login,
		Role:   user.Role,
	}

	token, err := s.generateJWT(&claims)
	if err != nil {
		return response, err
	}

	return models.SignUpResponse{
		JwtToken: token,
	}, err
}

func (s *Service) AdminSignIn(ctx context.Context, req models.AdminSignInRequest) (resp models.AdminSignInResponse, err error) {
	user, err := s.repo.SelectUserByLogin(ctx, req.Login)
	if err != nil {
		return resp, err
	}

	if !s.comparePassword(ctx, req.Password, user.Password) {
		return resp, errors.New("invalid creds")
	}

	claims := models.ClaimsJwt{
		UserID: user.ID,
		Login:  user.Login,
		Role:   user.Role,
	}

	token, err := s.generateJWT(&claims)
	if err != nil {
		return resp, err
	}

	return models.AdminSignInResponse{
		Token: token,
	}, err
}

func (s *Service) CreateInviteToken(ctx context.Context, req models.PostInviteTokenRequest) (resp models.PostInviteTokenResponse, err error) {
	token := entities.InviteToken{
		Content: uuid.New(),
		Role:    req.Role,
	}

	err = s.repo.InsertInviteToken(ctx, &token)
	if err != nil {
		return resp, err
	}

	return models.PostInviteTokenResponse{
		IviteToken: token.Content.String(),
	}, err
}

func (s *Service) hashPassword(ctx context.Context, password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func (s *Service) comparePassword(ctx context.Context, password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) generateJWT(claims *models.ClaimsJwt) (jwtToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err = token.SignedString([]byte(jwtHash))
	if err != nil {
		return jwtToken, err
	}

	return jwtToken, err
}
