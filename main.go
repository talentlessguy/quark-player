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

	slider := NewAudioBar(length.Seconds())

	go func() {
		for {

			slider.OnChanged = func(pos float64) {

				sliderPosToSeconds := int(pos / 100 * float64(length.Seconds()))

				fmt.Printf("Slider position in seconds %v\n", sliderPosToSeconds)
				fmt.Printf("Audio length in seconds %v\n", length.Seconds())
				fmt.Printf("Audio length in ms %v\n", length.Milliseconds())
				fmt.Printf("Streamer Length %v\n", ap.streamer.Len())

				audioPos := ap.streamer.Position()

				diff := sliderPosToSeconds - audioPos

				if diff > 0 {
					audioPos += diff
				} else {
					audioPos -= diff
				}

			}

			if !ap.ctrl.Paused {
				position = ap.sampleRate.D(ap.streamer.Position()).Round(time.Second)
				length = ap.sampleRate.D(ap.streamer.Len()).Round(time.Second)
				posLabel.SetText(DurationToString(position) + " / " + DurationToString(length))

				percentPos := (position.Seconds() / length.Seconds() * 100)

				slider.Value = percentPos

				slider.Refresh()

				posLabel.Refresh()
			}

		}
	}()

	controls := NewControls(ap)

	controlsGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(3), controls.prevTrackButton, controls.playButton, controls.nextTrackButton)

	sliderLayout := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), slider)

	grid := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), layout.NewSpacer(), image, layout.NewSpacer(), sliderLayout, layout.NewSpacer(), posLabel, label, NewSpacer(), controlsGrid)

	appContainer := fyne.NewContainerWithLayout(layout.NewCenterLayout(), grid)

	myWindow.SetContent(appContainer)
	myWindow.Resize(fyne.NewSize(400, 500))
	myWindow.ShowAndRun()
}
