package itunes

import (
	"fmt"
)

func (t *Track) isURLTrack() bool {
	return t.Kind() == ITTrackKindURL
}

// GetURL returns the URL of the stream represented by this track.
func (t *Track) GetURL() (string, error) {
	if !t.isURLTrack() {
		return "", fmt.Errorf("Not a URL track")
	}
	return t.getStringProperty("URL")
}

// SetURL sets the URL of the stream represented by this track.
func (t *Track) SetURL(url string) error {
	if !t.isURLTrack() {
		return fmt.Errorf("Not a URL track")
	}
	return t.setProperty("URL", url)
}
