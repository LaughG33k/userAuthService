package auth

import (
	"context"
	"errors"
	"time"

	"github.com/LaughG33k/userAuthService/iternal/model"
	"github.com/LaughG33k/userAuthService/pkg"
	"github.com/golang-jwt/jwt"
)

func (s *Service) UpdateSession(ctx context.Context, token model.TokenPair, fp model.FingerPrint) (model.TokenPair, error) {

	claims, err := s.jwtParse.ParseToken(token.Jwt)

	if err != nil {
		return model.TokenPair{}, err
	}

	session, err := s.sessionRepo.Get(ctx, token.Refresh)

	if err != nil {
		return model.TokenPair{}, err
	}

	if time.Now().Unix() > session.LifeTime {
		return model.TokenPair{}, errors.New("Token has expired")
	}

	jwt, err := s.jwtGen.NewToken(jwt.MapClaims{
		"uuid": claims["uuid"],
	})

	if err != nil {
		return model.TokenPair{}, err
	}

	refresh := pkg.GenerateRefreshToken(30)

	if session.Addr == fp.Addr && session.Browser == fp.Browser && session.Device == fp.Device {

		if err := s.createDeleteSession(ctx, model.Session{
			Token:    refresh,
			Owner:    claims["uuid"].(string),
			LifeTime: time.Now().Add(24 * time.Hour * 7).Unix(),
			FingerPrint: model.FingerPrint{
				Addr:    fp.Addr,
				Browser: fp.Browser,
				Device:  fp.Device,
			},
		}, token.Refresh); err != nil {
			return model.TokenPair{}, err
		}

		return model.TokenPair{Jwt: jwt, Refresh: refresh}, nil

	} else if session.Addr != fp.Addr && session.Browser == fp.Browser && session.Device == fp.Device {

		if err := s.createDeleteSession(ctx, model.Session{
			Token:    refresh,
			Owner:    claims["uuid"].(string),
			LifeTime: time.Now().Add(24 * time.Hour * 7).Unix(),
			FingerPrint: model.FingerPrint{
				Addr:    fp.Addr,
				Browser: fp.Browser,
				Device:  fp.Device,
			},
		}, token.Refresh); err != nil {
			return model.TokenPair{}, err
		}

		return model.TokenPair{Jwt: jwt, Refresh: refresh}, nil

	}

	if err := s.sessionRepo.Create(ctx, model.Session{
		Token:    refresh,
		Owner:    claims["uuid"].(string),
		LifeTime: time.Now().Add(24 * time.Hour * 7).Unix(),
		FingerPrint: model.FingerPrint{
			Addr:    fp.Addr,
			Browser: fp.Browser,
			Device:  fp.Device,
		},
	}); err != nil {
		return model.TokenPair{}, err
	}

	// to do send email warning

	return model.TokenPair{Jwt: jwt, Refresh: refresh}, nil
}

func (s *Service) createDeleteSession(ctx context.Context, new model.Session, oldToken string) error {

	if err := s.sessionRepo.Create(ctx, new); err != nil {
		return err
	}

	if err := s.sessionRepo.Delete(ctx, oldToken); err != nil {
		return err
	}

	return nil
}
