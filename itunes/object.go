package itunes

import (
	ole "github.com/go-ole/go-ole"
)

// COM is just an abstraction for *ole.IDispatch.
type COM struct {
	obj *ole.IDispatch
}

func (c *COM) IsNil() bool {
	return c.obj == nil
}

// ITunes is an object to abstract top level iTunes application. This represents
// IiTunes object in COM interface.
type ITunes struct {
	COM
}

// Playlist is an object to abstract the main library playlist. This is child
// for ITunes struct and represents IITPlaylist object in COM interface.
type Playlist struct {
	COM
}

// Playlist is an object to abstract the collection of library playlists. This is child
// for ITunes struct and represents IITPlaylistCollection object in COM interface.
type PlaylistCollection struct {
	COM
}

// TrackCollection is an object to abstract a collection of music track. This is child
// for Playlist struct and represents IITTrackCollection object in COM interface.
type TrackCollection struct {
	COM
}

// Track is an object to abstract a music track. This is child for Track struct
// and represents IITTrack object in COM interface.
type Track struct {
	COM
}
