package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	group, _ := errgroup.WithContext(ctx)
	group.Go(func() error {
		return serverSignal(ctx, cancel)
	})
	group.Go(func() error {
		return serverTCP(ctx, cancel)
	})

	if err := group.Wait(); err == nil {
		fmt.Println("Exit..")
	} else {
		fmt.Println("Err:", err)
	}
}

// serverSignal 监听信号
func serverSignal(ctx context.Context, cancel context.CancelFunc) error {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-ctx.Done():
		return nil
	case <-stop:
		cancel()
	}
	return nil
}

// serverTCP 启动TCP服务
func serverTCP(ctx context.Context, cancel context.CancelFunc) (err error) {
	listener, err := net.Listen("tcp", "localhost:18003")
	if err != nil {
		cancel()
		return err
	}
	// 开始监听 8000
	fmt.Println("tcp :18001")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.done close")
			_ = listener.Close()
			cancel()
			return nil
		default:
			accept, err := listener.Accept()
			if err != nil {
				_ = listener.Close()
				cancel()
				return err
			}

			ch := make(chan string, 8)
			go handleAccept(accept, ch)
			go sendMsg(accept, ch)
		}
	}
}

func handleAccept(c net.Conn, ch chan<- string) {
	defer c.Close()
	reader := bufio.NewReader(c)

	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		if err != nil {
			close(ch)
			return
		}

		ch <- string(line)
	}
}

func sendMsg(conn net.Conn, ch <-chan string) {
	wr := bufio.NewWriter(conn)

	for msg := range ch {
		_, _ = wr.WriteString(msg)
		_ = wr.Flush()
	}
}
