package App

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"os"
	"time"
)

type App struct{}

func (r App) Exit() {
	time.Sleep(time.Second * 10)
	os.Exit(1)
}

func (r App) GetImageWithURL(url string) (fyne.Resource, *canvas.Image) {
	image, err := fyne.LoadResourceFromURLString(url)
	if err != nil {
		return nil, nil
	}
	img := canvas.NewImageFromResource(image)
	img.FillMode = canvas.ImageFillContain
	return image, img
}
