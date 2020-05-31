package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// NewAudioBar - track bar
func NewAudioBar(totalTime float64) *widget.Slider {

	slider := widget.NewSlider(0, 100)

	slider.Resize(fyne.NewSize(300, theme.Padding()*2))

	slider.MinSize().Add(fyne.NewSize(300, theme.Padding()*2))

	return slider
}
