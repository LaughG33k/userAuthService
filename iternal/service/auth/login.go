package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/model"
	"github.com/LaughG33k/userAuthService/pkg"
	"github.com/golang-jwt/jwt"
)

func (s *Service) Login(ctx context.Context, user model.User, fp model.FingerPrint) (model.TokenPair, error) {

	uuid, err := s.userRepo.GetUuidByLP(ctx, user)

	if err != nil {
		return model.TokenPair{}, err
	}

	jwt, err := s.jwtGen.NewToken(jwt.MapClaims{
		"uuid": uuid,
	})

	if err != nil {
		return model.TokenPair{}, err
	}

	refresh := pkg.GenerateRefreshToken(30)

	if err := s.sessionRepo.Create(ctx, model.Session{
		Owner: uuid,
		Token: refresh,
		FingerPrint: model.FingerPrint{
			Addr:    fp.Addr,
			Browser: fp.Browser,
			Device:  fp.Device,
		},
	}); err != nil {
		return model.TokenPair{}, err
	}

	return model.TokenPair{
		Jwt:     jwt,
		Refresh: refresh,
	}, nil
}
