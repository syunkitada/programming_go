package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type App struct {
	loopInterval    time.Duration
	shutdownTimeout time.Duration
	isShutdown      bool
}

func NewApp() *App {
	app := App{
		loopInterval:    time.Duration(5) * time.Second,
		shutdownTimeout: time.Duration(10) * time.Second,
		isShutdown:      false,
	}
	return &app
}

func (app *App) Serv() error {
	for {
		fmt.Printf("%v: hello\n", time.Now())
		if app.isShutdown {
			fmt.Printf("Shutdown")
			os.Exit(0)
		}
		time.Sleep(app.loopInterval)
	}
	return nil
}

func (app *App) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
	defer cancel()
	app.isShutdown = true

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		os.Exit(1)
	}

	return nil
}

func main() {
	idleConnsClosed := make(chan struct{})
	app := NewApp()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(context.Background()); err != nil {
			fmt.Printf("App Shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	if err := app.Serv(); err != nil {
		fmt.Printf("App Serv Failed: %v\n", err)
	}
	<-idleConnsClosed
}
