package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//使用errGroup 进行错误管理goroutine间错误的记录
	g, ctx := errgroup.WithContext(context.Background())

	//创建路由对象
	mux := http.NewServeMux()
	// 创建url chan 控制goroutinue
	out := make(chan struct{})
	quit := make(chan os.Signal, 1)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		out <- struct{}{}
	})

	server :=http.Server{
		Handler: mux,
		Addr: ":8080",

	}
	// 任务1：开启服务器 任务1被任务2关闭
	g.Go(func() error {
		fmt.Println("Starting http server...")
		//go func() {
		//	//该进程
		//	<-ctx.Done()
		//	fmt.Println("ctx is done...")
		//
		//}()?? 为什么有时候在main goroutinue之后出现 有时候在之前出现
		return server.ListenAndServe()
	})

	// 任务2：服务器收到out或者sig退出
	g.Go(func() error {

		select {
		case <-out:
			server.Shutdown(ctx)  //shutdown之后立马出现http: Server closed
			fmt.Println("shutdown ")
			return errors.New("url simulation")
		case <-quit:
			return  errors.New("sig simulation")
		}

	} )

	// 任务3：信号
	g.Go(func() error {
		//无缓冲信号channel
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		// 无信号 则阻塞
		// 有信号 返回错误
		select {
		case <-ctx.Done():
			fmt.Println("ctx is over because of shutdown") //任务3比任务2fmt先打印，由于任务2fmt 后面？？
			return ctx.Err()
		}
	})
	// Wait blocks until all function calls from the Go method have returned, then
	// returns the first non-nil error (if any) from them.
	err := g.Wait()
	fmt.Println(err)
	fmt.Println("main goroutinue 退出")
}