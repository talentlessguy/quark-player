package main

import (
	"log"
	"os"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"

	"github.com/faiface/beep/mp3"
)

// AudioPanel - panel to control audio
type AudioPanel struct {
	sampleRate beep.SampleRate
	streamer   beep.StreamSeeker
	ctrl       *beep.Ctrl
	resampler  *beep.Resampler
	volume     *effects.Volume
}

func (ap *AudioPanel) play() {

	if ap.streamer.Position() == 0 {
		speaker.Play(ap.volume)
		ap.ctrl.Paused = false
	} else {
		ap.ctrl.Paused = false
		speaker.Unlock()
	}
}

func (ap *AudioPanel) pause() {
	if ap.streamer.Position() == ap.streamer.Len() {
		speaker.Close()

		ap.ctrl.Paused = true
	} else {
		speaker.Lock()
		ap.ctrl.Paused = true
	}
}

// NewAudioPanel - create new audio panel
func NewAudioPanel(sampleRate beep.SampleRate, streamer beep.StreamSeeker) *AudioPanel {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	return &AudioPanel{sampleRate, streamer, ctrl, resampler, volume}
}

// Audio - audio structure for decoding / getting data from it
type Audio struct {
	file *os.File
}

// NewAudio - init new audio from path
func NewAudio(path string) Audio {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return Audio{file: f}
}

// Decode - decode audio file
func (a *Audio) Decode() (s beep.StreamSeekCloser, format beep.Format, err error) {
	return mp3.Decode(a.file)
}

// ParseMetadata - return audio metadata
func (a *Audio) ParseMetadata() (m tag.Metadata) {
	m, err := tag.ReadFrom(a.file)

	if err != nil {
		log.Fatal(err)
	}

	return m
}
