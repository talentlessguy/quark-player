package main

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"github.com/dhowden/tag"
)

// LoadIcon - load icon image as fyne resource
func LoadIcon(iconPath string, iconName string) *fyne.StaticResource {
	iconFile, err := os.Open(iconPath)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(iconFile)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return fyne.NewStaticResource(iconName, b)
}

// GetImageFromMetadata - extract cover image from audio file
func GetImageFromMetadata(meta tag.Metadata) image.Image {
	img, _, err := image.Decode(bytes.NewReader(meta.Picture().Data))

	if err != nil {
		log.Fatal(err)
	}

	return img
}

// NewSpacer - create transparent rectangle
func NewSpacer() *canvas.Rectangle {
	bg := canvas.NewRectangle(color.Transparent)

	bg.Resize(fyne.NewSize(300, theme.Padding()*2))

	bg.SetMinSize(fyne.NewSize(300, theme.Padding()*2))

	return bg
}

// DurationToString - convert duration to MM:SS format
func DurationToString(time time.Duration) string {
	mins := time.Minutes()
	secs := time.Seconds()

	return strconv.Itoa(int(mins)) + ":" + strconv.Itoa(int(secs)-int(mins)*60)
}
