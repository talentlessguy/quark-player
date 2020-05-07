package main

import "os"

// GetAudioPath - get path to audio
func GetAudioPath() string {
	var path string

	if len(os.Args) <= 1 {
		path = "./test/Last Heroes - Dimensions.mp3"
	} else {
		path = os.Args[1]
	}

	return path

}
