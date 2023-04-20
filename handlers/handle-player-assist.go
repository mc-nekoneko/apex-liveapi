package handlers

import (
    "log"

    liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandlePlayerAssist(event *liveapi.LiveAPIEvent) {
    data := &liveapi.PlayerAssist{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}

