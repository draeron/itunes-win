// +build windows

// This program does various actions to iTunes.
package main

import (
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

	fo1, err := it.CreateFolder("test1")
	if err != nil || fo1 == nil {
		log.Fatal(err)
	}
	fo2, err := fo1.CreateFolder("folder")
	if err != nil || fo2 == nil {
		log.Fatal(err)
	}
	pl2, err := fo2.CreatePlaylist("playlist")
	if err != nil || pl2 == nil {
		log.Fatal(err)
	}
	err = pl2.SetParent(fo1)
	if err != nil {
		log.Fatal(err)
	}
}
