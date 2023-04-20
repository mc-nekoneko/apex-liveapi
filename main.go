package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	handlers "github.com/mc-nekoneko/apex-liveapi/handlers"
	liveapi "github.com/mc-nekoneko/apex-liveapi/liveapi"
	"google.golang.org/protobuf/proto"
)

var handlerMap map[string]func(*liveapi.LiveAPIEvent)

var upgrader = websocket.Upgrader{}

func handleWS(ginContext *gin.Context) {
	log.Println("Incoming request")
	log.Println(ginContext.Request)

	conn, err := upgrader.Upgrade(ginContext.Writer, ginContext.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorf("failed to read message: %v", err)
			return
		}

		event := &liveapi.LiveAPIEvent{}

		err = proto.Unmarshal(message, event)
		if err != nil {
			log.Errorf("failed to unmarshal message: %v", err)
			return
		}

		handler, ok := handlerMap[event.GameMessage.TypeUrl]
		if !ok {
			log.Errorf("no handler for message type %q", event.GameMessage.TypeUrl)
			return
		}

		handler(event)
	}
}

func init() {
	handlerMap = map[string]func(*liveapi.LiveAPIEvent){
		"rtech.liveapi.Vector3":                     handlers.HandleVector3,
		"rtech.liveapi.Player":                      handlers.HandlePlayer,
		"rtech.liveapi.CustomMatch_LobbyPlayer":     handlers.HandleCustomMatch_LobbyPlayer,
		"rtech.liveapi.Datacenter":                  handlers.HandleDatacenter,
		"rtech.liveapi.Version":                     handlers.HandleVersion,
		"rtech.liveapi.Init":                        handlers.HandleInit,
		"rtech.liveapi.CustomMatch_LobbyPlayers":    handlers.HandleCustomMatch_LobbyPlayers,
		"rtech.liveapi.ObserverSwitched":            handlers.HandleObserverSwitched,
		"rtech.liveapi.MatchSetup":                  handlers.HandleMatchSetup,
		"rtech.liveapi.GameStateChanged":            handlers.HandleGameStateChanged,
		"rtech.liveapi.CharacterSelected":           handlers.HandleCharacterSelected,
		"rtech.liveapi.MatchStateEnd":               handlers.HandleMatchStateEnd,
		"rtech.liveapi.RingStartClosing":            handlers.HandleRingStartClosing,
		"rtech.liveapi.RingFinishedClosing":         handlers.HandleRingFinishedClosing,
		"rtech.liveapi.PlayerConnected":             handlers.HandlePlayerConnected,
		"rtech.liveapi.PlayerDisconnected":          handlers.HandlePlayerDisconnected,
		"rtech.liveapi.PlayerStatChanged":           handlers.HandlePlayerStatChanged,
		"rtech.liveapi.PlayerDamaged":               handlers.HandlePlayerDamaged,
		"rtech.liveapi.PlayerKilled":                handlers.HandlePlayerKilled,
		"rtech.liveapi.PlayerDowned":                handlers.HandlePlayerDowned,
		"rtech.liveapi.PlayerAssist":                handlers.HandlePlayerAssist,
		"rtech.liveapi.SquadEliminated":             handlers.HandleSquadEliminated,
		"rtech.liveapi.GibraltarShieldAbsorbed":     handlers.HandleGibraltarShieldAbsorbed,
		"rtech.liveapi.PlayerRespawnTeam":           handlers.HandlePlayerRespawnTeam,
		"rtech.liveapi.PlayerRevive":                handlers.HandlePlayerRevive,
		"rtech.liveapi.ArenasItemSelected":          handlers.HandleArenasItemSelected,
		"rtech.liveapi.ArenasItemDeselected":        handlers.HandleArenasItemDeselected,
		"rtech.liveapi.InventoryPickUp":             handlers.HandleInventoryPickUp,
		"rtech.liveapi.InventoryDrop":               handlers.HandleInventoryDrop,
		"rtech.liveapi.InventoryUse":                handlers.HandleInventoryUse,
		"rtech.liveapi.BannerCollected":             handlers.HandleBannerCollected,
		"rtech.liveapi.PlayerAbilityUsed":           handlers.HandlePlayerAbilityUsed,
		"rtech.liveapi.ZiplineUsed":                 handlers.HandleZiplineUsed,
		"rtech.liveapi.GrenadeThrown":               handlers.HandleGrenadeThrown,
		"rtech.liveapi.BlackMarketAction":           handlers.HandleBlackMarketAction,
		"rtech.liveapi.WraithPortal":                handlers.HandleWraithPortal,
		"rtech.liveapi.AmmoUsed":                    handlers.HandleAmmoUsed,
		"rtech.liveapi.WeaponSwitched":              handlers.HandleWeaponSwitched,
		"rtech.liveapi.ChangeCamera":                handlers.HandleChangeCamera,
		"rtech.liveapi.PauseToggle":                 handlers.HandlePauseToggle,
		"rtech.liveapi.CustomMatch_CreateLobby":     handlers.HandleCustomMatch_CreateLobby,
		"rtech.liveapi.CustomMatch_JoinLobby":       handlers.HandleCustomMatch_JoinLobby,
		"rtech.liveapi.CustomMatch_LeaveLobby":      handlers.HandleCustomMatch_LeaveLobby,
		"rtech.liveapi.CustomMatch_SetReady":        handlers.HandleCustomMatch_SetReady,
		"rtech.liveapi.CustomMatch_GetLobbyPlayers": handlers.HandleCustomMatch_GetLobbyPlayers,
		"rtech.liveapi.CustomMatch_SetMatchmaking":  handlers.HandleCustomMatch_SetMatchmaking,
		"rtech.liveapi.CustomMatch_SetTeam":         handlers.HandleCustomMatch_SetTeam,
		"rtech.liveapi.CustomMatch_KickPlayer":      handlers.HandleCustomMatch_KickPlayer,
		"rtech.liveapi.CustomMatch_SetSettings":     handlers.HandleCustomMatch_SetSettings,
		"rtech.liveapi.CustomMatch_SendChat":        handlers.HandleCustomMatch_SendChat,
	}
}

func main() {
	server := gin.Default()
	server.GET("/", handleWS)
	server.Run(":7777")
}
