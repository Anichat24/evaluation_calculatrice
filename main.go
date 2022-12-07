package main

import (
	"evaluation/main.go/calculatrice"
	"evaluation/main.go/config"
	"evaluation/main.go/discord"
	"log"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		log.Fatal(err)
		return
	} else {
		discord.Run()
		<-make(chan struct{})
		return
	}
	calculatrice.Calculatrice(calcul []rune)
}
