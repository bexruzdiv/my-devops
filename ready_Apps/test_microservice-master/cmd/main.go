package main

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "genproto/sms_service"

	"bitbucket.org/alien_soft/sms_service/config"
	"bitbucket.org/alien_soft/sms_service/pkg/logger"
	sms "bitbucket.org/alien_soft/sms_service/pkg/sms"
	"bitbucket.org/alien_soft/sms_service/service"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "sms_service")
	defer logger.Cleanup(log)

	log.Info("main: pgxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	connDb := sqlx.MustConnect("postgres", psqlString)

	sendService := service.NewSendService(connDb)
	smsProviderService := service.NewSmsProviderService(connDb, log)
	smsPaymentsService := service.NewSmsPaymentsService(connDb, log)

	smsDaemon := sms.Daemon{Conf: &cfg, DB: connDb}
	go smsDaemon.Init()

	lis, err := net.Listen("tcp", cfg.RPCPort)

	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()

	pb.RegisterSmsServiceServer(s, sendService)
	pb.RegisterSmsProviderServiceServer(s, smsProviderService)
	pb.RegisterSmsPaymentsServiceServer(s, smsPaymentsService)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
