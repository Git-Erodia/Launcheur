package Mods

import (
	"erodialuncher/src/Zwuiix/App"
	"erodialuncher/src/Zwuiix/Game"
	"erodialuncher/src/Zwuiix/Minecraft"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	hook "github.com/robotn/gohook"
)

type Mods struct{}

var MenuOpen bool = false
var AppW fyne.Window
var ScaleButton *widget.Button

func (r Mods) Start() {
	s := hook.Start()
	defer hook.End()

	for {
		select {
		case i := <-s:
			if i.Kind > hook.KeyDown && i.Kind < hook.KeyUp {
				if i.Rawcode == 161 || i.Rawcode == 45 {
					if MenuOpen {
						AppW.Hide()
					} else {
						AppW.Show()
					}
					MenuOpen = !MenuOpen
				}
			}
		}
	}
}

func (r Mods) CreateApp(h *Game.Handler, app2 fyne.App) {
	AppW = app2.NewWindow("Érodia Launcheur 1.0.0 (" + Minecraft.Minecraft{}.GetMinecraftDID() + ")")

	icon, _ := App.App{}.GetImageWithURL("https://cdn.discordapp.com/attachments/1097233620118294688/1097233651936264292/logoerodiafondrouge.png")
	AppW.SetIcon(icon)

	label := widget.NewLabel("Voici l'interface des mods.")
	separator := widget.NewSeparator()
	separator.CreateRenderer()

	str := "Échelle de l'interface: "
	scale := widget.NewButton(str+"Normal", func() {
		if ScaleButton.Text == str+"Normal" {
			ScaleButton.SetText(str + "Petite")
			GuiScale{}.SetGuiScale(h, 7)
		} else if ScaleButton.Text == str+"Petite" {
			ScaleButton.SetText(str + "Normal")
			GuiScale{}.SetGuiScale(h, 3)
		}
	})
	ScaleButton = scale

	AppW.SetContent(container.NewVBox(label, separator, scale))
	AppW.SetCloseIntercept(func() {
		AppW.Hide()
		MenuOpen = false
	})
	AppW.Resize(fyne.NewSize(200, 200))
	AppW.SetFixedSize(true)
	AppW.SetMaster()
	AppW.CenterOnScreen()
}
