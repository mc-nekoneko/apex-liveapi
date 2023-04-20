package handlers

import (
    "log"

    liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandlePlayerStatChanged(event *liveapi.LiveAPIEvent) {
    data := &liveapi.PlayerStatChanged{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}

