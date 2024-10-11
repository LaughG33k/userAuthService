package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/converter"
	"github.com/LaughG33k/userAuthService/pkg/codegen"
)

func (a *AuthApi) Registration(ctx context.Context, req *codegen.RegReq) (*codegen.Empty, error) {

	if err := a.authService.Registration(ctx, converter.FromRegDescToUser(req)); err != nil {
		return nil, err
	}

	return nil, nil
}
