package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/converter"
	"github.com/LaughG33k/userAuthService/pkg/codegen"
)

func (a *AuthApi) Login(ctx context.Context, req *codegen.LoginReq) (*codegen.TokenPair, error) {

	tokenPair, err := a.authService.Login(ctx, req.Login, req.Password, converter.FromFPDescToFP(req.FingerPrint))

	if err != nil {
		return nil, err
	}

	return converter.FromTokenPairToTokenPairDesc(tokenPair), nil
}
