package itunes

import (
	"fmt"
)

func (t *Track) isURLTrack() bool {
	k, _ := getKind(t)
	if k == 3 {
		return true
	}
	return false
}

// GetURL returns the URL of the stream represented by this track.
func (t *Track) GetURL() (string, error) {
	if t.isURLTrack() == false {
		return "", fmt.Errorf("Not a URL track")
	}
	return getStringProperty(t, "URL")
}

// SetURL sets the URL of the stream represented by this track.
func (t *Track) SetURL(url string) error {
	if t.isURLTrack() == false {
		return fmt.Errorf("Not a URL track")
	}
	return setProperty(t, "URL", url)
}
