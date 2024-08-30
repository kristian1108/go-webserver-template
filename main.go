package main

import (
	"go-template/src/app"
	"log"
)

func main() {
	a := app.New()
	err := a.Run()
	if err != nil {
		log.Panicf("unable to start app, err:%s", err)
	}
}
