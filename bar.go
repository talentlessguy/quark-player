package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

func dragCircle(pos int) *canvas.Circle {
	circle := canvas.NewCircle(theme.PrimaryColor())

	circle.Resize(fyne.NewSize(20, 20))

	circle.MinSize().Add(fyne.NewSize(20, 20))

	circle.Move(fyne.NewPos(pos-5, -5))

	return circle
}

// NewAudioBar - track bar
func NewAudioBar(currentTime int, totalTime int) (*fyne.Container, *canvas.Rectangle, *canvas.Circle) {
	currentBar := canvas.NewRectangle(theme.PrimaryColor())

	currentBar.Resize(fyne.NewSize(currentTime, theme.Padding()*2))

	currentBar.SetMinSize(fyne.NewSize(currentTime, theme.Padding()*2))

	totalBar := canvas.NewRectangle(color.Black)

	totalBar.Resize(fyne.NewSize(300, theme.Padding()*2))

	totalBar.SetMinSize(fyne.NewSize(300, theme.Padding()*2))

	rngBtn := dragCircle(currentTime)

	c := fyne.NewContainer(totalBar, currentBar, rngBtn)

	cWithBg := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), NewSpacer(), c, NewSpacer())

	layout := fyne.NewContainerWithLayout(layout.NewCenterLayout(), cWithBg)

	return layout, currentBar, rngBtn
}
