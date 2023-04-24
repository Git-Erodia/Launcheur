package Zwuiix

import (
	_ "embed"
	"erodialuncher/src/Zwuiix/App"
	"erodialuncher/src/Zwuiix/Minecraft"
	"erodialuncher/src/Zwuiix/Minecraft/Mods"
	"erodialuncher/src/Zwuiix/Utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	golangdiscordrpc "github.com/SilverCory/golang_discord_rpc"
	"time"
)

var Hiden = false
var StartedGame = false
var StartedButton *widget.Button
var IsOpen = false

var Connection *golangdiscordrpc.RPCConnection

var VERSION = "1.0.4"

type Luncher struct {
}

func startDetect(a fyne.App, w fyne.Window) {
	go exit()

	Hiden = true
	w.Hide()

	time.Sleep(time.Second * 2)

	Minecraft.StartGame()
	StartedGame = true
	IsOpen = true

	time.Sleep(time.Second * 15)

	/*ip := Utils.Internet{}.GetIP()
	uuid := Minecraft.Minecraft{}.GetMinecraftDID()

	_, error := Utils.Request{}.WriteRequest(Utils.POST, uuid, ip)
	if error != nil {
		return
	}*/

	time.Sleep(time.Second * 15)

	Minecraft.Minecraft{}.StartGameHandler()

	Connection = golangdiscordrpc.NewRPCConnection("895391945722056775")
	err := Connection.Open()
	if err != nil {
	}
	_, err = Connection.Read()
	if err != nil {
	}

	Optimisation{}.StartOptimisation()

	Mods.Mods{}.CreateApp(Minecraft.Minecraft{}.GetGameHandler(), a)
	Mods.Mods{}.Start()
}

func exit() {
	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		if StartedGame && !IsOpen {
			App.App{}.Exit()
		}
	}
}

func (l Luncher) Start() {
	a := app.New()
	w := a.NewWindow("Érodia Launcheur " + VERSION)

	icon, _ := App.App{}.GetImageWithURL("http://cdn.erodia.fr/attachments/icon.png")
	w.SetIcon(icon)

	_, img := App.App{}.GetImageWithURL("http://cdn.erodia.fr/attachments/logo.png")
	img.Resize(fyne.NewSize(400, 300))

	start := widget.NewButton("Jouer", func() {
		if !Minecraft.IsGameOpen() {
			startDetect(a, w)
		}
	})
	StartedButton = start

	discord := widget.NewButton("Discord", func() { Utils.Link{}.Open("https://discord.gg/erodia") })
	startBox := container.NewVBox(start, discord)
	startBox.Move(fyne.NewPos(160, 230))

	go func() {
		ticker := time.NewTicker(time.Second * 5)

		for _ = range ticker.C {
			if !Hiden {
				if Minecraft.IsGameOpen() {
					StartedButton.Disable()
				} else {
					StartedButton.Enable()
				}
			}
		}
	}()

	w.SetContent(fyne.NewContainerWithoutLayout(img, startBox))

	w.Resize(fyne.NewSize(400, 380))
	w.SetFixedSize(true)
	w.SetMaster()
	w.CenterOnScreen()
	w.ShowAndRun()
}
