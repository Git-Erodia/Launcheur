package Minecraft

import (
	"bufio"
	"erodialuncher/src/Zwuiix/Game"
	"github.com/TheTitanrain/w32"
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"time"
)

type Minecraft struct{}

var MinecraftPath string = AppDataLocal + "Packages\\Microsoft.MinecraftUWP_8wekyb3d8bbwe\\"
var LocalState string = MinecraftPath + "LocalState\\"
var Games string = LocalState + "games\\"
var ComMojang string = Games + "com.mojang\\"
var MinecraftPe string = ComMojang + "minecraftpe\\"

var HS string = MinecraftPe + "hs"
var ClientID string = MinecraftPe + "clientId.txt"

var MinecraftDID string = ""
var MinecraftClientID string = ""

var GameHandler *Game.Handler = nil

func (r Minecraft) GetGameHandler() *Game.Handler {
	return GameHandler
}

func (r Minecraft) StartGameHandler() {
	GameHandler = Game.New()
}
func (r Minecraft) GetMinecraftDID() string {

	if MinecraftDID != "" {
		return MinecraftDID
	}

	readFile, err := os.Open(HS)

	if err != nil {
		return ""
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	did := ""
	i := 0

	for fileScanner.Scan() {
		if i == 1 {
			did = fileScanner.Text()
		} else {
			i++
		}
	}
	readFile.Close()

	MinecraftDID = did

	return MinecraftDID
}

func (r Minecraft) GetMinecraftClientID() string {
	if MinecraftClientID != "" {
		return MinecraftClientID
	}

	readFile, err := os.Open(ClientID)

	if err != nil {
		return ""
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	clientId := ""
	for fileScanner.Scan() {
		clientId = fileScanner.Text()
		break
	}
	readFile.Close()

	MinecraftClientID = clientId
	return MinecraftClientID
}

func (r Minecraft) GetUsername() string {
	return "Unknown"
}

func StartGame() {
	for IsGameOpen() == false {
		cmd := exec.Command("explorer.exe", "shell:appsFolder\\Microsoft.MinecraftUWP_8wekyb3d8bbwe!App")
		if err := cmd.Run(); err != nil {
			continue
		}

		for !IsGameOpen() {
			time.Sleep(time.Second * 1)
		}
		time.Sleep(time.Second * 1)
	}
}

func IsGameOpen() bool {

	open := func() bool {
		return w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft")) != 0
	}
	return open()
}
