package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/luofeng1/Go-001/Week04/api/user/v1"
	"github.com/luofeng1/Go-001/Week04/internal/service"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// init service
	us := InitUser()
	s := service.NewUserService(us)

	// register grpc service
	gRPCServer := grpc.NewServer()
	pb.RegisterUserServer(gRPCServer, s)

	// test
	reflection.Register(gRPCServer)

	ctx, cancel := context.WithCancel(context.TODO())
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return serverGRPC(ctx, gRPCServer)
	})
	g.Go(func() error {
		return serverSignal(ctx, cancel)
	})
	if err := g.Wait(); err == nil {
		fmt.Println("Exit..")
	} else {
		fmt.Println("Err:", err)
	}
}

// serverGRPC 启动gRPC
func serverGRPC(ctx context.Context, gRPCServer *grpc.Server) (err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("stop grpc server")
			gRPCServer.GracefulStop()
		}
	}()
	listener, err := net.Listen("tcp", ":8880")
	fmt.Println("grpc: 8880")
	return gRPCServer.Serve(listener)
}

// serverSignal 监听信号
func serverSignal(ctx context.Context, cancel context.CancelFunc) error {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	select {
	case <-ctx.Done():
		return nil
	case <-stop:
		cancel()
	}
	return nil
}
