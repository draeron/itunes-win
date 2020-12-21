// +build windows

// This program does various actions to main library. It plays first track,
// set shuffle on and pause. Next it scans and searches tracks of specified
// artist and album ("In Utero" by Nirvana in this example).
package main

import (
	"fmt"
	"log"

	"github.com/draeron/itunes-win/itunes"
)

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
	// Search the album "In Utero" by Nirvana. Print track names if found.
	count, tracks, err := p.SearchArtists("Nirvana")
	if err != nil {
		log.Fatal(err)
	}
	var ret string
	for i := 1; i <= count; i++ {
		track, _ := tracks.GetTrackByIndex(i)
		album, _ := track.GetAlbum()
		if album != "In Utero" {
			// It could be "Bleach", "Nevermind" and so on.
			continue
		}
		name, _ := track.GetName()
		ret += fmt.Sprintf("%s\n", name)
	}
	fmt.Printf(ret)
}
