package main

import (
	"bufio"
	"bytes"
	"image"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne"
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
