package handlers

import (
    "log"

    liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandleCustomMatch_LobbyPlayer(event *liveapi.LiveAPIEvent) {
    data := &liveapi.CustomMatch_LobbyPlayer{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}

