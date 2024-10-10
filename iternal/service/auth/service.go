package auth

import (
	"github.com/LaughG33k/userAuthService/iternal/repository"
	"github.com/LaughG33k/userAuthService/iternal/service"
	"github.com/LaughG33k/userAuthService/pkg"
)

type Service struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	jwtGen      pkg.JwtGenerator
	jwtParse    pkg.JwtParser
}

func New(userRepo repository.UserRepository, sessionRepo repository.SessionRepository, jwtGen pkg.JwtGenerator, jwtParse pkg.JwtParser) service.Auth {
	return &Service{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}
