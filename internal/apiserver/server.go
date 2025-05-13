package apiserver

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/srxstack/gin-template/internal/pkg/log"
	genericoptions "github.com/srxstack/srxstack/pkg/options"
	"github.com/srxstack/srxstack/pkg/server"
)

type Config struct {
	ServerMode  string
	JWTKey      string
	Expiration  time.Duration
	HTTPOptions *genericoptions.HTTPOptions
	TLSOptions  *genericoptions.TLSOptions
}

type UnionServer struct {
	srv server.Server
}

type ServerConfig struct {
	cfg *Config
}

func (cfg *Config) NewUnionServer() (*UnionServer, error) {
	serverConfig, err := cfg.NewServerConfig()
	if err != nil {
		return nil, err
	}

	log.Infow("Initializing federation server", "server-mode", cfg.ServerMode)

	srv, err := serverConfig.NewGinServer(), nil
	if err != nil {
		return nil, err
	}

	return &UnionServer{srv: srv}, nil
}

func (s *UnionServer) Run() error {
	go s.srv.RunOrDie()

	// 创建一个 os.Signal 类型的 channel，用于接收系统信号
	quit := make(chan os.Signal, 1)
	// 当执行 kill 命令时（不带参数），默认会发送 syscall.SIGTERM 信号
	// 使用 kill -2 命令会发送 syscall.SIGINT 信号（例如按 CTRL+C 触发）
	// 使用 kill -9 命令会发送 syscall.SIGKILL 信号，但 SIGKILL 信号无法被捕获，因此无需监听和处理
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞程序，等待从 quit channel 中接收到信号
	<-quit

	log.Infow("Shutting down server ...")

	// 优雅关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 先关闭依赖的服务，再关闭被依赖的服务
	s.srv.GracefulStop(ctx)

	log.Infow("Server exited")
	return nil
}

func (cfg *Config) NewServerConfig() (*ServerConfig, error) {
	return &ServerConfig{cfg: cfg}, nil
}
