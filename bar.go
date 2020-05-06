package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

// NewAudioBar - track bar
func NewAudioBar(currentTime int, timeLeft int) *fyne.Container {
	currentBar := canvas.NewRectangle(theme.PrimaryColor())

	totalBar := canvas.NewRectangle(color.Black)

	currentBar.Resize(fyne.NewSize(100, theme.Padding()*2))

	totalBar.Resize(fyne.NewSize(300, theme.Padding()*2))

	c := fyne.NewContainer(totalBar, currentBar)

	spacer := canvas.NewRectangle(color.Transparent)

	spacer.Resize(fyne.NewSize(1000, theme.Padding()*2))

	layout := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), spacer, c, spacer)

	return layout
}
