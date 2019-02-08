package main

import "github.com/Liar233/empty-app/internal/app/definition"

func main() {
	application := definition.InitApp()
	application.Start()
}
