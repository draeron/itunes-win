package itunes

import (
	"fmt"
	"os"
)

// Delete deletes this track.
func (t *Track) Delete() error {
	if _, err := t.obj.CallMethod("Delete"); err != nil {
		return err
	}
	return nil
}

// Play plays this track.
func (t *Track) Play() error {
	if _, err := t.obj.CallMethod("Play"); err != nil {
		return err
	}
	return nil
}

// AddArtwork adds artwork from an image file to this track.
func (t *Track) AddArtwork(file string) error {
	_, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("%v: image file access error", err)
	}
	if _, err := t.obj.CallMethod("AddArtworkFromFile", file); err != nil {
		return err
	}
	return nil
}

func getStringProperty(track *Track, property string) (string, error) {
	n, err := track.obj.GetProperty(property)
	if err != nil {
		return "", err
	}
	return n.Value().(string), nil
}

func getInt64Property(track *Track, property string) (int64, error) {
	n, err := track.obj.GetProperty(property)
	if err != nil {
		return 0, err
	}
	return n.Val, nil
}

func getBoolProperty(track *Track, property string) (bool, error) {
	n, err := track.obj.GetProperty(property)
	if err != nil {
		return false, err
	}
	return n.Value().(bool), nil
}

func getKind(track *Track) (int64, error) {
	return getInt64Property(track, "Kind")
}

// GetKind returns the kind of the track in string.
func (t *Track) GetKind() (string, error) {
	var ret string
	k, err := getKind(t)
	switch k {
	case 0:
		ret = "Unknown"
	case 1:
		ret = "File track"
	case 2:
		ret = "CD track"
	case 3:
		ret = "URL track"
	case 4:
		ret = "Device track"
	case 5:
		ret = "Shared library track"
	}
	return ret, err
}

// GetAlbum returns the name of the album containing the track.
func (t *Track) GetAlbum() (string, error) {
	return getStringProperty(t, "Album")
}

// GetArtist returns the name of the artist of the track.
func (t *Track) GetArtist() (string, error) {
	return getStringProperty(t, "Artist")
}

// GetBitRate returns the bit rate of the track (in kbps).
func (t *Track) GetBitRate() (int64, error) {
	return getInt64Property(t, "BitRate")
}

// GetBPM returns the BPM of the track.
func (t *Track) GetBPM() (int64, error) {
	return getInt64Property(t, "BPM")
}

// GetComment returns the comment of the track.
func (t *Track) GetComment() (string, error) {
	return getStringProperty(t, "Comment")
}

// GetComposer returns the composer of the track.
func (t *Track) GetComposer() (string, error) {
	return getStringProperty(t, "Composer")
}

// GetDuration returns the length of the track.
func (t *Track) GetDuration() (int64, error) {
	return getInt64Property(t, "Duration")
}

// GetGenre returns the genre of the track.
func (t *Track) GetGenre() (string, error) {
	return getStringProperty(t, "Genre")
}

// GetName returns the name of the track.
func (t *Track) GetName() (string, error) {
	return getStringProperty(t, "Name")
}

// GetPlayedCount returns the number of times the track has been played.
func (t *Track) GetPlayedCount() (int64, error) {
	return getInt64Property(t, "PlayedCount")
}

// GetSize returns the size of the track (in bytes).
func (t *Track) GetSize() (int64, error) {
	return getInt64Property(t, "Size")
}

// GetTrackNumber returns the index of the track on album.
func (t *Track) GetTrackNumber() (int64, error) {
	return getInt64Property(t, "TrackNumber")
}

// GetYear returns the year the track was recorded/released.
func (t *Track) GetYear() (int64, error) {
	return getInt64Property(t, "Year")
}

// GetCategory returns the category for the track.
func (t *Track) GetCategory() (string, error) {
	return getStringProperty(t, "Category")
}

// GetDescription returns the description for the track.
func (t *Track) GetDescription() (string, error) {
	return getStringProperty(t, "Description")
}

// GetAlbumRating returns the user or computed rating of the album that this
// track belongs to (0 to 100). If the album rating has never been set, or has
// been set to 0.
func (t *Track) GetAlbumRating() (int64, error) {
	return getInt64Property(t, "AlbumRating")
}

func setProperty(t *Track, property string, val interface{}) error {
	_, err := t.obj.PutProperty(property, val)
	if err != nil {
		return err
	}
	return nil
}

// SetAlbum sets the name of the album containing the track.
func (t *Track) SetAlbum(val string) error {
	return setProperty(t, "Album", val)
}

// SetArtist sets the name of the artist of the track.
func (t *Track) SetArtist(val string) error {
	return setProperty(t, "Artist", val)
}

// SetBPM sets the BPM of the track.
func (t *Track) SetBPM(val int) error {
	return setProperty(t, "BPM", val)
}

// SetComment sets the comment of the track.
func (t *Track) SetComment(val string) error {
	return setProperty(t, "Comment", val)
}

// SetComposer sets the composer of the track.
func (t *Track) SetComposer(val string) error {
	return setProperty(t, "Composer", val)
}

// SetGenre sets the genre of the track.
func (t *Track) SetGenre(val string) error {
	return setProperty(t, "Genre", val)
}

// SetName sets the name of the track.
func (t *Track) SetName(val string) error {
	return setProperty(t, "Name", val)
}

// SetPlayedCount sets the number of times the track has been played.
func (t *Track) SetPlayedCount(val int) error {
	return setProperty(t, "PlayedCount", val)
}

// SetTrackNumber sets the index of the track on album.
func (t *Track) SetTrackNumber(val int) error {
	return setProperty(t, "TrackNumber", val)
}

// SetYear sets the year the track was recorded/released.
func (t *Track) SetYear(val int) error {
	return setProperty(t, "Year", val)
}

// SetCategory sets the category for the track.
func (t *Track) SetCategory(val string) error {
	return setProperty(t, "Category", val)
}

// SetDescription the description for the track.
func (t *Track) SetDescription(val string) error {
	return setProperty(t, "Description", val)
}

// SetRememberBookmark sets whether the playback position is remembered for this
// track.
func (t *Track) SetRememberBookmark(val bool) error {
	return setProperty(t, "RememberBookmark", val)
}

// SetAlbumRating sets the user or computed rating of the album that this
// track belongs to (0 to 100). If the album rating is set to 0, it will be
// computed based on the ratings of tracks in the album.
func (t *Track) SetAlbumRating(val int) error {
	return setProperty(t, "AlbumRating", val)
}
