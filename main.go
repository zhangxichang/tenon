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

var upgrader = websocket.Upgrader{}

func handle_ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("WebSocket连接升级错误: %v\n", err)
		return
	}
	defer conn.Close()
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("消息读取错误: %v\n", err)
			return
		}
		fmt.Println(string(msg))
		if err := conn.WriteMessage(mt, []byte("我收到了")); err != nil {
			fmt.Printf("消息发送错误: %v\n", err)
			return
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("build")))
	mux.HandleFunc("/ws", handle_ws)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	sig_chan := make(chan os.Signal, 1)
	signal.Notify(sig_chan, os.Interrupt, syscall.SIGTERM)
	go func() {
		fmt.Println("WebUI: http://127.0.0.1:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("服务错误: %v\n", err)
		}
	}()
	<-sig_chan
	fmt.Println("服务关闭中...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("服务关闭时发生错误: %v\n", err)
		os.Exit(-1)
	}
	fmt.Println("服务已关闭")
}
