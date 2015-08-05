// +build windows

// This program does various actions to iTunes.
package main

import (
	"log"

	"github.com/tdsh/itunes-win/itunes"
)

func main() {
	// Initialization.
	it, err := itunes.Init()
	// Make sure do Exit() at the end.
	defer it.Exit()
	if err != nil {
		log.Fatal(err)
	}
	// Now you can control iTunes via ITunes object.
	// In the below example error handling is ignored for simplification.

	// Play current track.
	it.Play()
	// Mute.
	it.Mute(true)
	// Go to the next track.
	it.NextTrack()
	// Set mute off.
	it.Mute(false)
	// Set volume 50.
	it.SoundVolume(50)
	// Pause the track play.
	it.Pause()
	// Quit iTunes application.
	it.QuitApp()
}
