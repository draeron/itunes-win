/*
Package itunes-win enables you to handle iTunes on Windows in Go. This is
a thin wrapper of iTunes COM interface.

Note this library doesn't relect all of the methods and the properties of
iTunes COM interface.

Initialization must be done first of all by Init. Exit also must be called
at the end:

	it, err := itunes.Init()
	defer it.Exit()

For control over iTunes application, use ITunes object returned:

	it.Play()

For control over main library playlist, get Playlist via GetMainPlaylist:

	p, err := it.GetMainPlaylist()
	if err != nil {
		// handle error
	}
	p.PlayFirstTrack()

For control over tracks in the playlist, get Tracks via GetTracks. Likewise,
you can control each track in the tracks via Track got by GetTrackByIndex or
GetTrackByName:

	t, err := p.GetTracks()
	// ...
	c64, _ := tracks.GetCount()
	count := int(c64)
	for i := 1; i <= count; i++ {
		track, _ := tracks.GetTrackByIndex(i)
	}

*/
package itunes // import "github.com/draeron/itunes-win/itunes"
