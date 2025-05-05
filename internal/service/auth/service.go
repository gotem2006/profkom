package auth

import (
	"profkom/internal/repository/auth"

	txmanager "github.com/avito-tech/go-transaction-manager/trm/manager"
)

const (
	jwtHash = "asdqwe2131241eqeqw"
)

type Service struct {
	repo      *auth.Repository
	txManager *txmanager.Manager
}

func New(repo *auth.Repository, txManager *txmanager.Manager) *Service {
	return &Service{
		repo:      repo,
		txManager: txManager,
	}
}
