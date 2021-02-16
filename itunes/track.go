package itunes

import (
	"fmt"
	"os"
	"time"
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

// GetKind returns the kind of the track in string.
func (t *Track) Kind() ITTrackKind {
	return ITTrackKind(t.getKind())
}

func (t *Track) GetPersistentIDLow() (int64, error) {
	return t.getInt64Property("ITObjectPersistentIDLow")
}

func (t *Track) GetPersistentIDHigh() (int64, error) {
	return t.getInt64Property("ITObjectPersistentIDHigh")
}

// GetAlbum returns the name of the album containing the track.
func (t *Track) GetAlbum() (string, error) {
	return t.getStringProperty("Album")
}

// GetArtist returns the name of the artist of the track.
func (t *Track) GetArtist() (string, error) {
	return t.getStringProperty("Artist")
}

// GetBitRate returns the bit rate of the track (in kbps).
func (t *Track) GetBitRate() (int64, error) {
	return t.getInt64Property("BitRate")
}

// GetBPM returns the BPM of the track.
func (t *Track) GetBPM() (int64, error) {
	return t.getInt64Property("BPM")
}

// GetComment returns the comment of the track.
func (t *Track) GetComment() (string, error) {
	return t.getStringProperty("Comment")
}

// GetComposer returns the composer of the track.
func (t *Track) GetComposer() (string, error) {
	return t.getStringProperty("Composer")
}

// GetDuration returns the length of the track.
func (t *Track) GetDuration() (int64, error) {
	return t.getInt64Property("Duration")
}

// GetGenre returns the genre of the track.
func (t *Track) GetGenre() (string, error) {
	return t.getStringProperty("Genre")
}

// GetName returns the name of the track.
func (t *Track) GetName() (string, error) {
	return t.getStringProperty("Name")
}

// GetPlayedCount returns the number of times the track has been played.
func (t *Track) GetPlayedCount() (int64, error) {
	return t.getInt64Property("PlayedCount")
}

// GetSize returns the size of the track (in bytes).
func (t *Track) GetSize() (int64, error) {
	return t.getInt64Property("Size")
}

// GetTrackNumber returns the index of the track on album.
func (t *Track) GetTrackNumber() (int64, error) {
	return t.getInt64Property("TrackNumber")
}

// GetYear returns the year the track was recorded/released.
func (t *Track) GetYear() (int64, error) {
	return t.getInt64Property("Year")
}

// GetCategory returns the category for the track.
func (t *Track) GetCategory() (string, error) {
	return t.getStringProperty("Category")
}

// GetDescription returns the description for the track.
func (t *Track) GetDescription() (string, error) {
	return t.getStringProperty("Description")
}

func (t *Track) GetRating() (int64, error) {
	return t.getInt64Property("Rating")
}

func (t *Track) GetDateAdded() (time.Time, error) {
	return t.getDateProperty("DateAdded")
}

// GetAlbumRating returns the user or computed rating of the album that this
// track belongs to (0 to 100). If the album rating has never been set, or has
// been set to 0.
func (t *Track) GetAlbumRating() (int64, error) {
	return t.getInt64Property("AlbumRating")
}

// SetAlbum sets the name of the album containing the track.
func (t *Track) SetAlbum(val string) error {
	return t.setProperty("Album", val)
}

// SetArtist sets the name of the artist of the track.
func (t *Track) SetArtist(val string) error {
	return t.setProperty("Artist", val)
}

// SetBPM sets the BPM of the track.
func (t *Track) SetBPM(val int) error {
	return t.setProperty("BPM", val)
}

// SetComment sets the comment of the track.
func (t *Track) SetComment(val string) error {
	return t.setProperty("Comment", val)
}

// SetComposer sets the composer of the track.
func (t *Track) SetComposer(val string) error {
	return t.setProperty("Composer", val)
}

// SetGenre sets the genre of the track.
func (t *Track) SetGenre(val string) error {
	return t.setProperty("Genre", val)
}

// SetName sets the name of the track.
func (t *Track) SetName(val string) error {
	return t.setProperty("Name", val)
}

// SetPlayedCount sets the number of times the track has been played.
func (t *Track) SetPlayedCount(val int) error {
	return t.setProperty("PlayedCount", val)
}

// SetTrackNumber sets the index of the track on album.
func (t *Track) SetTrackNumber(val int) error {
	return t.setProperty("TrackNumber", val)
}

// SetYear sets the year the track was recorded/released.
func (t *Track) SetYear(val int) error {
	return t.setProperty("Year", val)
}

// SetCategory sets the category for the track.
func (t *Track) SetCategory(val string) error {
	return t.setProperty("Category", val)
}

// SetDescription the description for the track.
func (t *Track) SetDescription(val string) error {
	return t.setProperty("Description", val)
}

// SetRememberBookmark sets whether the playback position is remembered for this
// track.
func (t *Track) SetRememberBookmark(val bool) error {
	return t.setProperty("RememberBookmark", val)
}

func (t *Track) SetRating(val int) error {
	return t.setProperty("Rating", val)
}

// SetAlbumRating sets the user or computed rating of the album that this
// track belongs to (0 to 100). If the album rating is set to 0, it will be
// computed based on the ratings of tracks in the album.
func (t *Track) SetAlbumRating(val int) error {
	return t.setProperty("AlbumRating", val)
}
