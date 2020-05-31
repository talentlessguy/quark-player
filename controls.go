package main

import "fyne.io/fyne/widget"

// Controls -  set of player controls
type Controls struct {
	playButton      *widget.Button
	prevTrackButton *widget.Button
	nextTrackButton *widget.Button
}

// NewControls - return a set of player controls
func NewControls(ap *AudioPanel) Controls {

	return Controls{
		playButton:      NewPlayButton(ap),
		prevTrackButton: NewPrevTrackButton(ap),
		nextTrackButton: NewNextTrackButton(),
	}
}
