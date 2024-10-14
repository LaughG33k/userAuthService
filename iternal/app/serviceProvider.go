package app

import (
	"context"
	"log"

	"github.com/LaughG33k/userAuthService/client/postgresql"
	"github.com/LaughG33k/userAuthService/iternal/config"
	"github.com/LaughG33k/userAuthService/iternal/repository"
	"github.com/LaughG33k/userAuthService/iternal/repository/session"
	usrRepo "github.com/LaughG33k/userAuthService/iternal/repository/user"
	"github.com/LaughG33k/userAuthService/iternal/service"
	"github.com/LaughG33k/userAuthService/iternal/service/auth"
	"github.com/LaughG33k/userAuthService/pkg"
	"github.com/golang-jwt/jwt"
)

type serviceProvider struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
	authService       service.Auth
	jwtWorker         *pkg.JwtWorker
	cfg               config.AppConfig
}

func newServiceProvider(config config.AppConfig) *serviceProvider {
	return &serviceProvider{
		cfg: config,
	}
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {

	if s.userRepository != nil {
		return s.userRepository
	}

	client, err := postgresql.NewClient(ctx, 3, s.cfg.UserDb)

	if err != nil {
		log.Panic(err)
	}

	s.userRepository = usrRepo.NewUserRepostiroy(client)

	return s.userRepository
}

func (s *serviceProvider) JwtWorker(_ context.Context) *pkg.JwtWorker {

	if s.jwtWorker != nil {
		return s.jwtWorker
	}

	w, err := pkg.NewJwtWorker([]byte("secret-key"), jwt.SigningMethodES256)

	if err != nil {
		log.Panic(err)
	}

	s.jwtWorker = w

	return s.jwtWorker
}

func (s *serviceProvider) SessionRepository(ctx context.Context) repository.SessionRepository {

	if s.sessionRepository != nil {
		return s.sessionRepository
	}

	client, err := postgresql.NewClient(ctx, 3, s.cfg.SessionDb)

	if err != nil {
		log.Panic(err)
	}

	s.sessionRepository = session.NewSessionRepository(client)

	return s.sessionRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.Auth {

	if s.authService != nil {
		return s.authService
	}

	s.authService = auth.New(s.UserRepository(ctx), s.SessionRepository(ctx), s.JwtWorker(ctx), s.JwtWorker(ctx))

	return s.authService

}
