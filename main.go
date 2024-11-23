package main

import (
	"log"

	"github.com/mehmetymw/thundio/configs"
)

func main() {
	_, err := configs.NewConfig()
	if err != nil {
		log.Fatalln("config cannot create")
		return
	}

}
