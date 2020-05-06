package main

import (
	"log"
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/faiface/beep/speaker"
)

func main() {
	myApp := app.New()

	path := os.Args[1]

	progress := NewAudioBar(10, 100)

	audio := NewAudio(path)

	meta := audio.ParseMetadata()

	title := meta.Artist() + " - " + meta.Title()

	myWindow := myApp.NewWindow("Quark Player")
	text := widget.NewLabel(title)

	text.Alignment = fyne.TextAlignCenter

	streamer, format, err := audio.Decode()

	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ap := NewAudioPanel(format.SampleRate, streamer)

	controls := NewControls(ap)

	image := &canvas.Image{
		FillMode: canvas.ImageFillContain,
		Image:    GetImageFromMetadata(meta),
	}

	image.SetMinSize(fyne.Size{Width: 300, Height: 300})
	controlsGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(3), controls.prevTrackButton, controls.playButton, controls.nextTrackButton)

	grid := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), layout.NewSpacer(), image, layout.NewSpacer(), progress, layout.NewSpacer(), text, controlsGrid)

	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(400, 500))
	myWindow.ShowAndRun()
}
