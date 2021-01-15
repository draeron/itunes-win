package itunes

import (
	"fmt"
)

func (t *Track) isFileOrCDTrack() bool {
	k, _ := getKind(t)
	if k == 1 || k == 2 {
		return true
	}
	return false
}

// GetLocation returns the full path to the file represented by this track.
func (t *Track) GetLocation() (string, error) {
	if t.isFileOrCDTrack() == false {
		return "", fmt.Errorf("Not a file or CD track")
	}
	return getStringProperty(t, "Location")
}

func (t *Track) SetLocation(location string) error {
	if t.isFileOrCDTrack() == false {
		return fmt.Errorf("Not a file or CD track")
	}
	return setStringProperty(t, "Location", location)
}

// GetRememberBookmark returns true if playback position is remembered for this
// track.
func (t *Track) GetRememberBookmark() (bool, error) {
	if t.isFileOrCDTrack() == false {
		return false, fmt.Errorf("Not a file or CD track")
	}
	return getBoolProperty(t, "RememberBookmark")
}

// GetExcludeFromShuffle true if this track is skipped when shuffling.
func (t *Track) GetExcludeFromShuffle() (bool, error) {
	if t.isFileOrCDTrack() == false {
		return false, fmt.Errorf("Not a file or CD track")
	}
	return getBoolProperty(t, "ExcludeFromShuffle")
}

// GetLylics returns the lyrics for the track.
func (t *Track) GetLylics() (string, error) {
	if t.isFileOrCDTrack() == false {
		return "", fmt.Errorf("Not a file or CD track")
	}
	return getStringProperty(t, "Lylics")
}

// GetAlbumArtist returns the name of the album artist of the track.
func (t *Track) GetAlbumArtist() (string, error) {
	if t.isFileOrCDTrack() == false {
		return "", fmt.Errorf("Not a file or CD track")
	}
	return getStringProperty(t, "AlbumArtist")
}
