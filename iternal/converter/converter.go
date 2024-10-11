package converter

import (
	"github.com/LaughG33k/userAuthService/iternal/model"
	"github.com/LaughG33k/userAuthService/pkg/codegen"
)

func FromRegDescToUser(reg *codegen.RegReq) model.User {
	return model.User{
		Login:    reg.Login,
		Password: reg.Password,
		Email:    reg.Email,
		Name:     reg.Name,
	}
}

func FromTokenPairDescToTokenPair(pair *codegen.TokenPair) model.TokenPair {

	return model.TokenPair{
		Jwt:     pair.Jwt,
		Refresh: pair.RefreshToken,
	}
}

func FromFPDescToFP(fp *codegen.FingerPrint) model.FingerPrint {
	return model.FingerPrint{
		Addr:    fp.Addr,
		Browser: fp.Browser,
		Device:  fp.Device,
	}
}

func FromTokenPairToTokenPairDesc(pair model.TokenPair) *codegen.TokenPair {
	return &codegen.TokenPair{
		Jwt:          pair.Jwt,
		RefreshToken: pair.Refresh,
	}
}
