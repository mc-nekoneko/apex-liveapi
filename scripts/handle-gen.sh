#!/bin/zsh

while read -r event; do
  package=$(echo $event | cut -d'.' -f1-2)
  event_name_def=$(echo $event | cut -d'.' -f3)

  package=$(echo $package | tr '[:upper:]' '[:lower:]')
  event_name=$(echo $event_name_def | sed 's/\([A-Z]\)/-\L\1/g' | sed 's/^-//' | tr '_' '-')

  filename="handle-${event_name}.go"

  echo "package handlers

import (
    \"log\"

    liveapi \"github.com/mc-nekoneko/apex-liveapi/liveapi\"
)

func Handle${event_name_def}(event *liveapi.LiveAPIEvent) {
    data := &liveapi.${event_name_def}{}
    err := event.GameMessage.UnmarshalTo(data)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}
" > "${filename}"

done < events.txt
