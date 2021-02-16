package itunes

import (
	"fmt"
	"os"
)

const (
	// searchFieldAll searches all fields of each track.
	searchFieldAll = iota
	// searchFieldVisible searches only the fields with columns that are
	// currently visible in the display for the playlist.
	searchFieldVisible
	// searchFieldArtists searches only the artist field of each track.
	searchFieldArtists
	// searchFieldAlbums searches only the album field of each track.
	searchFieldAlbums
	// searchFieldComposers searches only the composer field of each track.
	searchFieldComposers
	// searchFieldSongNames searches only the song name field of each track.
	searchFieldSongNames
)

// Delete deletes this playlist.
func (p *Playlist) Delete() error {
	if _, err := p.obj.CallMethod("Delete"); err != nil {
		return err
	}
	return nil
}

// PlayFirstTrack starts play the first track in this playlist.
func (p *Playlist) PlayFirstTrack() error {
	if _, err := p.obj.CallMethod("PlayFirstTrack"); err != nil {
		return err
	}
	return nil
}

func (p *Playlist) Kind() ITPlaylistKind {
	return ITPlaylistKind(p.getKind())
}

func search(p *Playlist, word string, field int) (int, *TrackCollection, error) {
	r, err := p.obj.CallMethod("Search", word, field)
	if err != nil {
		return -1, nil, err
	}
	o := COM{
		obj: r.ToIDispatch(),
	}
	t := TrackCollection{COM: o}
	r, _ = t.obj.GetProperty("Count")
	count := int(r.Val)
	return count, &t, nil
}

// SearchAll searches all fields of each track.
// It returns the number of track found, tracks found as *Tracks and error if any.
func (p *Playlist) SearchAll(word string) (int, *TrackCollection, error) {
	return search(p, word, searchFieldAll)
}

// SearchVisible searches only the fields with columns that are currently visible
// in the display for the playlist.
// It returns the number of track found, tracks found as *Tracks and error if any.
func (p *Playlist) SearchVisible(word string) (int, *TrackCollection, error) {
	return search(p, word, searchFieldVisible)
}

// SearchArtists searches only the artist field of each track.
// It returns the number of track found, tracks found as *Tracks and error if any.
func (p *Playlist) SearchArtists(word string) (int, *TrackCollection, error) {
	return search(p, word, searchFieldArtists)
}

// SearchAlbums searches only the album field of each track.
// It returns the number of track found, tracks found as *Tracks and error if any.
func (p *Playlist) SearchAlbums(word string) (int, *TrackCollection, error) {
	return search(p, word, searchFieldAlbums)
}

// SearchComposers searches only the composer field of each track.
// It returns the number of track found, tracks found as *Tracks and error if any.
func (p *Playlist) SearchComposers(word string) (int, *TrackCollection, error) {
	return search(p, word, searchFieldComposers)
}

// SearchSongNames searches only the song name field of each track.
// It returns the number of track found, tracks found as *Tracks and error if any.
func (p *Playlist) SearchSongNames(word string) (int, *TrackCollection, error) {
	return search(p, word, searchFieldSongNames)
}

// AddFile adds path to this playlist.
func (p *Playlist) AddFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("iTunes library directory access error\n\n%v", err)
	}
	_, err = p.obj.CallMethod("AddFile", path)
	if err != nil {
		return err
	}
	return nil
}

// GetTracks returns Tracks in this playlist.
func (p *Playlist) GetTracks() (*TrackCollection, error) {
	r, err := p.obj.GetProperty("Tracks")
	if err != nil {
		return nil, err
	}
	o := COM{
		obj: r.ToIDispatch(),
	}
	t := TrackCollection{COM: o}
	return &t, nil
}

// SetShuffle sets this playlist shuffle mode (played in random order).
func (p *Playlist) SetShuffle(val bool) error {
	if _, err := p.obj.PutProperty("Shuffle", val); err != nil {
		return err
	}
	return nil
}

const (
	// RepeatOff sets repeat mode off.
	RepeatOff = iota
	// RepeatSong repeats the song.
	RepeatSong
	// RepeatAll repeats the playlist.
	RepeatAll
)

// SetRepeat sets playlist repeat mode.
// Please specify any one of RepeatOff, RepeatSong or RepeatAll.
func (p *Playlist) SetRepeat(mode int) error {
	if _, err := p.obj.PutProperty("SongRepeat", mode); err != nil {
		return err
	}
	return nil
}
