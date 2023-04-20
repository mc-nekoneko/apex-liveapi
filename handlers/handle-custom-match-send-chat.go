package handlers

import (
	"log"

	liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandleCustomMatch_SendChat(event *liveapi.LiveAPIEvent) {
	data := &liveapi.CustomMatch_SendChat{}
	err := event.GameMessage.UnmarshalTo(data)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(data)
}
