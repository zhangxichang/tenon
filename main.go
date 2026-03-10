package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	ctx, exit := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer exit()
	devMode := len(os.Args) > 1 && os.Args[1] == "--dev"
	mux := http.NewServeMux()
	if !devMode {
		mux.Handle("/", http.FileServer(http.Dir("dist")))
	}
	mux.HandleFunc("/ws", handle_ws)
	server := http.Server{
		Addr:    "localhost:10270",
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("服务出错: %v\n", err)
			os.Exit(-1)
		}
	}()
	if devMode {
		fmt.Println("开发模式, WebSocket: ws://localhost:10270/ws")
	} else {
		fmt.Println("WebUI: http://localhost:10270")
	}
	fmt.Println("按 Ctrl+C 退出")
	<-ctx.Done()
	ctx, close_server := context.WithTimeout(context.Background(), 5*time.Second)
	defer close_server()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("服务关闭时出错: %v\n", err)
		os.Exit(-1)
	}
}

func handle_ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("WebSocket连接升级错误: %v\n", err)
		os.Exit(-1)
	}
	defer conn.Close()
	if err := handle_conn(conn); err != nil {
		fmt.Printf("连接处理错误: %v\n", err)
		os.Exit(-1)
	}
}

func handle_conn(conn *websocket.Conn) error {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		fmt.Println(string(msg))
	}
}
