package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/converter"
	"github.com/LaughG33k/userAuthService/pkg/codegen"
)

func (a *AuthApi) Logout(ctx context.Context, req *codegen.TokenPair) (*codegen.Empty, error) {

	if err := a.authService.Logout(ctx, converter.FromTokenPairDescToTokenPair(req)); err != nil {
		return nil, err
	}

	return nil, nil
}
