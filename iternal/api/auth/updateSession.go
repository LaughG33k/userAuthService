package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/converter"
	"github.com/LaughG33k/userAuthService/pkg/codegen"
)

func (a *AuthApi) UpdateSession(ctx context.Context, req *codegen.TokenPairForUpdate) (*codegen.TokenPair, error) {

	tokenPair, err := a.authService.UpdateSession(ctx, converter.FromTokenPairDescToTokenPair(req.Pair), converter.FromFPDescToFP(req.FingerPrint))

	if err != nil {
		return nil, err
	}

	return converter.FromTokenPairToTokenPairDesc(tokenPair), nil
}
