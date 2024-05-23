package app

import (
	"github.com/LaughG33k/userAuthService/iternal"
	"github.com/LaughG33k/userAuthService/iternal/dbClient/postgresql"
	grpcserver "github.com/LaughG33k/userAuthService/iternal/grpcServer"
	"github.com/LaughG33k/userAuthService/iternal/handler"
	"github.com/LaughG33k/userAuthService/iternal/repository"

	"context"
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AppConfig struct {
	UserDB      DBCSettings        `json:"user_db"`
	GrpcServer  GrpcServerSettings `json:"grpc_server"`
	ServiceUuid string             `json:"service_uuid"`
}

type DBCSettings struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

type GrpcServerSettings struct {
	Addr         string `json:"addr"`
	MaxConcConns int    `json:"max_concurrency_connections"`
}

func Run() {

	cfg, err := os.ReadFile("/Users/user/Desktop/docker_user_vol/cfg/userAuthService/config.json")
	logger := initLogger()

	if logger == nil {
		return
	}

	if err != nil {
		logger.Error(err.Error())
		return
	}

	var appCfg AppConfig

	if err := json.Unmarshal(cfg, &appCfg); err != nil {
		logger.Error(err.Error())
		return
	}

	ctx := context.Background()
	userDb, err := postgresql.NewClient(ctx, 5, appCfg.UserDB.Name, appCfg.UserDB.Password, appCfg.UserDB.Host, appCfg.UserDB.Port, appCfg.UserDB.Db)

	if err != nil {
		logger.Error(err.Error())
		return
	}

	userRepository := repository.NewUserRepostiroy(ctx, userDb)
	rtRepository := repository.NewRefreshTokenRepostiroy(ctx, userDb)

	jwtWorker := iternal.NewJwtWorker(appCfg.ServiceUuid)
	authHandler := handler.NewAuthHandler(userRepository, rtRepository, jwtWorker)

	server := grpcserver.NewServer(
		ctx,
		logger,
		appCfg.GrpcServer.Addr,
		appCfg.GrpcServer.MaxConcConns,
		authHandler,
	)

	fmt.Println("start")

	if err = server.Start(); err != nil {
		logger.Error(err.Error())
	}

}

func initLogger() *zap.Logger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, err := os.OpenFile("/Users/user/Desktop/messengerMicroservice/github.com/LaughG33k/userAuthService/log.json", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Panic(err)
		return nil
	}

	writer := zapcore.AddSync(logFile)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger

}
