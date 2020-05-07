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

	currentBar.Resize(fyne.NewSize(100, theme.Padding()*2))

	currentBar.SetMinSize(fyne.NewSize(100, theme.Padding()*2))

	totalBar := canvas.NewRectangle(color.Black)

	totalBar.Resize(fyne.NewSize(300, theme.Padding()*2))

	totalBar.SetMinSize(fyne.NewSize(300, theme.Padding()*2))

	c := fyne.NewContainer(totalBar, currentBar)

	cWithBg := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), NewSpacer(), c, NewSpacer())

	layout := fyne.NewContainerWithLayout(layout.NewCenterLayout(), cWithBg)

	return layout
}
