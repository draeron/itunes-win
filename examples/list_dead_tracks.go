// +build windows

// This program scans main library and lists dead tracks whose files are not
// accessible or no longer available.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/draeron/itunes-win/itunes"
)

// Please change it to an appropriate location.
const logFile = "C:/Users/jdoe/Desktop/dead_tracks.log"

func main() {
	// Initialization.
	it, err := itunes.Init()
	// Make sure do Exit() at the end.
	defer it.Exit()
	if err != nil {
		log.Fatal(err)
	}
	// Get main library.
	p, _ := it.GetMainPlaylist()
	// Now you can control main library via Playlist object.
	// In the below example error handling is ignored for simplification.

	tracks, _ := p.GetTracks()
	c64, _ := tracks.GetCount()
	count := int(c64)
	result := make([]byte, 0)
	// It can take long time to scan library if it has many tracks.
	for i := 1; i <= count; i++ {
		track, _ := tracks.GetTrackByIndex(i)
		path, _ := track.GetLocation()
		// Location is empty if it's a dead track.
		if path != "" {
			continue
		}
		name, _ := track.GetName()
		album, _ := track.GetAlbum()
		artist, _ := track.GetArtist()
		result = append(result, fmt.Sprintf("%s (%s) by %s\r\n", name, album, artist)...)
	}
	f, _ := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	f.Write(result)
	f.Close()
}
