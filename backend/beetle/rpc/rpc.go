package rpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ShintaNakama/yukibu-channel/backend/beetle/env"
	"github.com/ShintaNakama/yukibu-channel/backend/beetle/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Register service register
type Register func(svr *grpc.Server)

// RPCServer gRPC server
type RPCServer struct {
	name string
	port int
}

// RPCServerOption サーバーの起動オプション
type RPCServerOption struct {
	Name string
	Port int
}

var log *zap.Logger

// NewRPCServer returns RPCServer
func NewRPCServer(option *RPCServerOption) *RPCServer {
	return &RPCServer{
		name: option.Name,
		port: option.Port,
	}
}

// Run start server
func (sv *RPCServer) Run(register Register) {
	time.Local, _ = time.LoadLocation("Asia/Tokyo")
	e := env.NewENV()
	log = logger.NewLogger(e)
	defer func() {
		_ = log.Sync()
	}()

	server := sv.newServer()
	register(server)
	health.RegisterHealthServer(server, &healthServer{})
	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", sv.port))
	if err != nil {
		log.Fatal("failed to listen", zap.Int("port", sv.port))
	}
	defer listen.Close()

	go func() {
		if err := server.Serve(listen); err != nil {
			zap.L().Error("error occurred", zap.Error(err))
		}
	}()

	zap.L().Info("server start")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	zap.L().Info(fmt.Sprintf("SIGNAL %d received, then shutdown...", <-quit))

	server.GracefulStop()
	zap.L().Info("server shutdown")
}

func (sv *RPCServer) newServer() *grpc.Server {
	return grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			bypassHealthCheckLog(),
		),
	)
}

func bypassHealthCheckLog() grpc.UnaryServerInterceptor {
	loggerInterceptor := grpc_zap.UnaryServerInterceptor(log)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod == "/grpc.health.v1.Health/Check" {
			return handler(ctx, req)
		}
		return loggerInterceptor(ctx, req, info, handler)
	}
}

type healthServer struct {
	health.HealthServer
}

func (h *healthServer) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (h *healthServer) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return nil
}
