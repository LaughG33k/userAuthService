package handler

import (
	"context"
	"time"

	jwt "github.com/LaughG33k/userAuthService/iternal"
	"github.com/LaughG33k/userAuthService/iternal/repository"
	"github.com/LaughG33k/userAuthService/pkg"
	"github.com/LaughG33k/userAuthService/pkg/grpc/codegen/authservice/authservice/codegen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcAuthHandler struct {
	codegen.UnimplementedAuthServer
	OperationTimeout time.Duration
	userRepository   *repository.UserRepository
	rtRepository     *repository.RefreshTokenRepository
	jwtWorker        *jwt.JwtWorker
}

func NewAuthHandler(userRepository *repository.UserRepository, refreshTokenRepository *repository.RefreshTokenRepository, jwtWorker *jwt.JwtWorker) *GrpcAuthHandler {

	return &GrpcAuthHandler{
		userRepository: userRepository,
		rtRepository:   refreshTokenRepository,
		jwtWorker:      jwtWorker,
	}

}

func (h *GrpcAuthHandler) Registration(ctx context.Context, in *codegen.RegReq) (*codegen.RegResp, error) {

	if len(in.Login) > 30 || len(in.Password) > 30 || len(in.Name) > 30 || len(in.Email) > 256 {
		return nil, status.Errorf(codes.InvalidArgument, "the length of the entered data exceeds the possible length")
	}

	tm, canc := context.WithTimeout(ctx, h.OperationTimeout)

	defer canc()

	err := h.userRepository.CreateUser(tm, in.Login, in.Password, in.Name, in.Email)

	if err != nil {

		if err.Error() == "23505" {
			return nil, status.Errorf(codes.AlreadyExists, "")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &codegen.RegResp{}, nil
}
func (h *GrpcAuthHandler) Login(ctx context.Context, in *codegen.LoginReq) (*codegen.LoginResp, error) {

	if len(in.Login) > 30 || len(in.Password) > 30 {
		return nil, status.Errorf(codes.InvalidArgument, "the length of the entered data exceeds the possible length")
	}

	tm, canc := context.WithTimeout(ctx, h.OperationTimeout)

	defer canc()

	uuid, err := h.userRepository.CheckUserByLP(tm, in.Login, in.Password)

	if err != nil {

		if err.Error() == "no rows in result set" {
			return nil, status.Error(codes.NotFound, "")
		}

		return nil, status.Error(codes.Internal, "")
	}

	rt := pkg.GenerateRefreshToken(20)

	err = h.rtRepository.CreateRefreshToken(tm, "", rt, uuid, time.Now().Add(24*time.Hour*15).Unix())

	if err != nil {
		return nil, status.Error(codes.Internal, "server error")
	}

	token, err := h.jwtWorker.CreateJwt(
		map[string]any{
			"uuid": uuid,
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "server error")
	}

	return &codegen.LoginResp{Jwt: token, RefreshToken: rt}, nil
}
func (h *GrpcAuthHandler) UpdateJwt(ctx context.Context, in *codegen.UpdateJwtReq) (*codegen.UpdateJwtResp, error) {

	tm, canc := context.WithTimeout(ctx, h.OperationTimeout)

	defer canc()

	uuid, timeLife, err := h.rtRepository.FindRefreshToken(tm, in.RefreshToken)

	if err != nil {

		if err.Error() == "no rows in result set" {
			return nil, status.Error(codes.NotFound, "")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	if time.Now().Unix() > timeLife {
		return nil, status.Error(codes.PermissionDenied, "token is not valid")
	}

	rt := pkg.GenerateRefreshToken(20)

	if err = h.rtRepository.CreateRefreshToken(tm, in.RefreshToken, rt, uuid, time.Now().Add(24*time.Hour*15).Unix()); err != nil {
		return nil, status.Errorf(codes.Internal, "server error")
	}

	token, err := h.jwtWorker.CreateJwt(map[string]any{
		"uuid": uuid,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "server error")
	}

	return &codegen.UpdateJwtResp{Jwt: token, RefreshToken: rt}, nil
}
