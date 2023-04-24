package Zwuiix

import (
	"erodialuncher/src/Zwuiix/Minecraft"
	"erodialuncher/src/Zwuiix/RPC"
	"time"
)

type Optimisation struct{}

func (r Optimisation) StartOptimisation() {
	go updateGame()
	go rpc()
}

func updateGame() {
	ticker := time.NewTicker(time.Second * 3)

	for _ = range ticker.C {
		if StartedGame == true && Minecraft.IsGameOpen() == true {
			IsOpen = true
		} else {
			IsOpen = false
		}
	}
}

func rpc() {
	ticker := time.NewTicker(time.Second * 30)

	for _ = range ticker.C {
		if StartedGame && IsOpen {
			RPC.RPCD{}.Update(Connection)
		}
	}
}
