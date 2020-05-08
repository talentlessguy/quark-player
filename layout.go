package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/dhowden/tag"
)

// NewTrackTitleLabel - get track title label element
func NewTrackTitleLabel(title string) *widget.Label {
	text := widget.NewLabel(title)

	text.Alignment = fyne.TextAlignCenter

	return text
}

// NewCoverImage - create image block using metadata
func NewCoverImage(meta tag.Metadata) *canvas.Image {
	image := &canvas.Image{
		FillMode: canvas.ImageFillContain,
		Image:    GetImageFromMetadata(meta),
	}

	image.SetMinSize(fyne.NewSize(300, 300))

	return image

}

// NewTimeInfo - layout with two labels representing track time info
func NewTimeInfo(pos time.Duration, len time.Duration) *widget.Label {

	str := pos.String() + " / " + len.String()

	posLabel := widget.NewLabel(str)

	posLabel.Alignment = fyne.TextAlignCenter

	return posLabel
}
