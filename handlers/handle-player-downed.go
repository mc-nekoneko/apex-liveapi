package handlers

import (
    "log"

    liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandlePlayerDowned(event *liveapi.LiveAPIEvent) {
    data := &liveapi.PlayerDowned{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}

