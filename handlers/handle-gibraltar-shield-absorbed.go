package handlers

import (
    "log"

    liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
)

func HandleGibraltarShieldAbsorbed(event *liveapi.LiveAPIEvent) {
    data := &liveapi.GibraltarShieldAbsorbed{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}

