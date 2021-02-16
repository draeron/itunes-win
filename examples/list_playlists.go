// +build windows

// This program does various actions to iTunes.
package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/dustin/go-humanize"

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

	lists, err := it.Playlists()
	if err != nil {
		log.Fatal(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 4, 0, ' ', tabwriter.Debug)
	defer writer.Flush()
	writer.Write([]byte("Name \t Full Path \t Kind \t Duration \t Bytes Size \t Special Kind \t Smart\n"))

	for _, list := range lists {
		smart, _ := list.IsSmart()
		size, _ := list.Size()
		duration, _ := list.Duration()
		skind, _ := list.SpecialKind()

		path := ""
		parent, _ := list.Parent()
		for parent != nil {
			path += "/" + parent.Name()
			parent, _ = parent.Parent()
		}
		path += "/" + list.Name()

		writer.Write([]byte(fmt.Sprintf(" %v \t %v \t %v \t %v \t %v \t %v \t %v \n", list.Name(), path, list.Kind(), duration, humanize.Bytes(size), skind, smart)))
		// fmt.Printf("Kind: %v, Path: %v, Size: %v, Smart: %v\n", list.Kind(), path + "/" + list.Name(), size, smart)
	}
}
