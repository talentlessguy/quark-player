package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func getIcons(iconName string) *fyne.StaticResource {
	iconSet := map[string]*fyne.StaticResource{
		"play":  LoadIcon("icons/play-fill.png", "play"),
		"pause": LoadIcon("icons/pause-fill.png", "pause"),
		"prev":  LoadIcon("icons/skip-back-fill.png", "prev"),
		"next":  LoadIcon("icons/skip-forward-fill.png", "next"),
	}

	return iconSet[iconName]
}

// NewPlayButton - create play / pause button
func NewPlayButton(ap *AudioPanel) *widget.Button {
	btn := &widget.Button{

		Icon: getIcons("play"),
	}

	btn.OnTapped = func() {

		if btn.Icon.Name() == "play" {
			ap.play()
			btn.Icon = getIcons("pause")
		} else {
			ap.pause()
			btn.Icon = getIcons("play")
		}
	}

	return btn
}

// NewPrevTrackButton - run previous track
func NewPrevTrackButton() *widget.Button {
	btn := widget.NewButtonWithIcon("", getIcons("prev"), func() {
		fmt.Println("Prev Track")
	})

	return btn
}

// NewNextTrackButton - run previous track
func NewNextTrackButton() *widget.Button {
	btn := widget.NewButtonWithIcon("", getIcons("next"), func() {
		fmt.Println("Next Track")
	})

	return btn
}
