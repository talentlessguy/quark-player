package main

import (
	"log"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"github.com/faiface/beep/speaker"
)

func main() {
	myApp := app.New()

	path := GetAudioPath()

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

	posLabel := NewTimeInfo(position, length)

	progress, currentBar, circle := NewAudioBar(int(position.Seconds()), int(length.Seconds()))

	go func() {
		for {

			if !ap.ctrl.Paused {
				position = ap.sampleRate.D(ap.streamer.Position()).Round(time.Second)
				length = ap.sampleRate.D(ap.streamer.Len()).Round(time.Second)
				posLabel.SetText(DurationToString(position) + " / " + DurationToString(length))

				percentPos := int(position.Seconds() / length.Seconds() * 100)

				barSize := fyne.NewSize(percentPos, theme.Padding()*2)

				currentBar.Resize(barSize)

				currentBar.SetMinSize(barSize)

				circle.Move(fyne.NewPos(percentPos-5, -5))
				posLabel.Refresh()
				currentBar.Refresh()
			}

		}
	}()

	controls := NewControls(ap)

	controlsGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(3), controls.prevTrackButton, controls.playButton, controls.nextTrackButton)

	grid := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), layout.NewSpacer(), image, layout.NewSpacer(), progress, layout.NewSpacer(), posLabel, label, NewSpacer(), controlsGrid)

	appContainer := fyne.NewContainerWithLayout(layout.NewCenterLayout(), grid)

	myWindow.SetContent(appContainer)
	myWindow.Resize(fyne.NewSize(400, 500))
	myWindow.ShowAndRun()
}
