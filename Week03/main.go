/**
package main
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出
*/
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return serverHTTP(ctx)
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

func serverHTTP(ctx context.Context) (err error) {
	s := http.Server{}
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("stop http server")
			_ = s.Shutdown(context.Background())
		}
	}()
	if err = s.ListenAndServe(); err == http.ErrServerClosed {
		return nil
	}
	return err
}

func serverSignal(ctx context.Context, cancel context.CancelFunc) error {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		return nil
	case <-stop:
		cancel()
	}
	return nil
}
