package itunes

import (
	"fmt"
)

func (t *Track) isFileOrCDTrack() bool {
	kind := t.Kind()
	if kind == ITTrackKindFile || kind == ITTrackKindCD {
		return true
	}
	return false
}

// GetLocation returns the full path to the file represented by this track.
func (t *Track) GetLocation() (string, error) {
	if t.isFileOrCDTrack() == false {
		return "", fmt.Errorf("Not a file or CD track")
	}
	return t.getStringProperty("Location")
}

func (t *Track) SetLocation(location string) error {
	if t.isFileOrCDTrack() == false {
		return fmt.Errorf("Not a file or CD track")
	}
	return t.setStringProperty("Location", location)
}

// GetRememberBookmark returns true if playback position is remembered for this
// track.
func (t *Track) GetRememberBookmark() (bool, error) {
	if t.isFileOrCDTrack() == false {
		return false, fmt.Errorf("Not a file or CD track")
	}
	return t.getBoolProperty("RememberBookmark")
}

// GetExcludeFromShuffle true if this track is skipped when shuffling.
func (t *Track) GetExcludeFromShuffle() (bool, error) {
	if t.isFileOrCDTrack() == false {
		return false, fmt.Errorf("Not a file or CD track")
	}
	return t.getBoolProperty("ExcludeFromShuffle")
}

// GetLylics returns the lyrics for the track.
func (t *Track) GetLylics() (string, error) {
	if t.isFileOrCDTrack() == false {
		return "", fmt.Errorf("Not a file or CD track")
	}
	return t.getStringProperty("Lylics")
}

// GetAlbumArtist returns the name of the album artist of the track.
func (t *Track) GetAlbumArtist() (string, error) {
	if t.isFileOrCDTrack() == false {
		return "", fmt.Errorf("Not a file or CD track")
	}
	return t.getStringProperty("AlbumArtist")
}
