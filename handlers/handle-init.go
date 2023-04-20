package handlers

import (
	"log"

	liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandleInit(event *liveapi.LiveAPIEvent) {
	data := &liveapi.Init{}
	err := event.GameMessage.UnmarshalTo(data)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(data)
}
