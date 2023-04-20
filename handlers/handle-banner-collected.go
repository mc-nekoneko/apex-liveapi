package handlers

import (
    "log"

    liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandleBannerCollected(event *liveapi.LiveAPIEvent) {
    data := &liveapi.BannerCollected{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}

