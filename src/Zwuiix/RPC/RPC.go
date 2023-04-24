package RPC

import (
	"encoding/json"
	"erodialuncher/src/Zwuiix/Minecraft"
	"fmt"
	golangdiscordrpc "github.com/SilverCory/golang_discord_rpc"
	"os"
)

type RPCD struct{}

func (r RPCD) Update(Connection *golangdiscordrpc.RPCConnection) {

	fmt.Println("[RICH-PRESENCE] Update...")
	username := Minecraft.Minecraft{}.GetUsername()

	name := "Joueur non connecté"
	if username != "NoConnect" {
		name = username
	}

	str := "Dans le menu"
	if name != "Joueur non connecté" {
		str = "Serveur : Faction"
	}

	activity := &golangdiscordrpc.Activity{
		Details:  "Serveur PVP/Faction",
		State:    str,
		Instance: true,
		Assets: &golangdiscordrpc.Assets{
			LargeText:    "erodia.fr",
			LargeImageID: "logoerodiafondrouge",
			SmallText:    name,
			SmallImageID: "steve",
		},
		TimeStamps: &golangdiscordrpc.TimeStamps{
			EndTimestamp:   1,
			StartTimestamp: 1,
		},
	}

	presence := &golangdiscordrpc.CommandRichPresenceMessage{
		CommandMessage: golangdiscordrpc.CommandMessage{Command: "SET_ACTIVITY"},
		Args: &golangdiscordrpc.RichPresenceMessageArgs{
			Pid:      os.Getpid(),
			Activity: activity,
		},
	}

	presence.SetNonce()
	data, err := json.Marshal(presence)

	if err != nil {
		return
	}

	err = Connection.Write(string(data))
	if err != nil {
		return
	}

	_, err = Connection.Read()
	if err != nil {
		return
	}

	return
}
