package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/akthrmsx/todo-app-go/config"
	"github.com/akthrmsx/todo-app-go/router"
	"github.com/akthrmsx/todo-app-go/server"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(0)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with %v", url)

	r, err := router.NewRouter(ctx, cfg)
	if err != nil {
		return err
	}

	svr := server.NewServer(l, r)
	return svr.Run(ctx)
}
