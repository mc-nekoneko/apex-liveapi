package handlers

import (
    "log"

    liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandlePlayerDisconnected(event *liveapi.LiveAPIEvent) {
    data := &liveapi.PlayerDisconnected{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}

