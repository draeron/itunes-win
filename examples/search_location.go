// +build windows

// This program lists file location of tracks of specified
// artist and album ("Run The Jewels 3" by Run The Jewels in this example).
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tdsh/itunes-win/itunes"
)

const logFile = "C:/Users/jdoe/Desktop/filelist.log"

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

	// Play the first track in p.
	p.PlayFirstTrack()
	// Set this playlist shffule mode on.
	p.SetShuffle(true)
	// Pause the play.
	it.Pause()
	// Search the album "Run The Jewels 3" by Run The Jewels. write track names if found.
	count, tracks, err := p.SearchArtists("Run The Jewels")
	if err != nil {
		log.Fatal(err)
	}
	result := make([]byte, 0)
	for i := 1; i <= count; i++ {
		track, _ := tracks.GetTrackByIndex(i)
		album, _ := track.GetAlbum()
		if album != "Run The Jewels 3" {
			// Skip the other albums.
			continue
		}
		name, _ := track.GetName()
		path, _ := track.GetLocation()
		result = append(result, fmt.Sprintf("%s: %s\r\n", name, path)...)
	}
	f, _ := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	f.Write(result)
	f.Close()
}
