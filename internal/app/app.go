package app

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

type Application struct {
	stop chan bool
}

func (app *Application) Start() {
	signs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signs
		fmt.Println()
		fmt.Println("Stopping...")
		done <- true
	}()

	<-done

	fmt.Println("Exit")
}

func NewApplication() Application {
	return Application{
		stop: make(chan bool),
	}
}
