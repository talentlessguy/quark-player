package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"github.com/faiface/beep/speaker"
)

func main() {
	myApp := app.New()

	path := GetAudioPath()

	progress := NewAudioBar(10, 100)

	audio := NewAudio(path)

	meta := audio.ParseMetadata()

	title := meta.Artist() + " - " + meta.Title()

	myWindow := myApp.NewWindow("Quark Player")

	label := NewTrackTitleLabel(title)

	image := NewCoverImage(meta)

	streamer, format, err := audio.Decode()

	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ap := NewAudioPanel(format.SampleRate, streamer)

	var position time.Duration
	var length time.Duration

	timeInfo, posLabel, lenLabel := NewTimeInfo(position, length)

	go func() {
		for {

			position = ap.sampleRate.D(ap.streamer.Position()).Round(time.Second)
			length = ap.sampleRate.D(ap.streamer.Len()).Round(time.Second)
			posLabel.SetText(position.String())
			lenLabel.SetText(length.String())
			fmt.Printf("%v of %v", position, length)
			fmt.Println()
			timeInfo.Refresh()
		}
	}()

	controls := NewControls(ap)

	controlsGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(3), controls.prevTrackButton, controls.playButton, controls.nextTrackButton)

	grid := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), layout.NewSpacer(), image, layout.NewSpacer(), progress, layout.NewSpacer(), timeInfo, label, NewSpacer(), controlsGrid)

	appContainer := fyne.NewContainerWithLayout(layout.NewCenterLayout(), grid)

	myWindow.SetContent(appContainer)
	myWindow.Resize(fyne.NewSize(400, 500))
	myWindow.ShowAndRun()
}
