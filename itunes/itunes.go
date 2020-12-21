package itunes

import (
	"fmt"
	"net/url"

	"github.com/go-ole/go-ole"
)

// GetMainPlaylist returns main library playlist, Playlist.
func (i *ITunes) GetMainPlaylist() (*Playlist, error) {
	r, err := i.obj.GetProperty("LibraryPlaylist")
	if err != nil {
		return nil, err
	}
	o := COM{
		obj: r.ToIDispatch(),
	}
	p := Playlist{COM: o}
	return &p, nil
}

// BackTrack repositions to the beginning of the current track or go to the
// previous track if already at start of current track.
func (i *ITunes) BackTrack() error {
	if _, err := i.obj.CallMethod("BackTrack"); err != nil {
		return err
	}
	return nil
}

// NextTrack advances to the next track in the current playlist.
func (i *ITunes) NextTrack() error {
	if _, err := i.obj.CallMethod("NextTrack"); err != nil {
		return err
	}
	return nil
}

// Pause pauses current track play.
func (i *ITunes) Pause() error {
	if _, err := i.obj.CallMethod("Pause"); err != nil {
		return err
	}
	return nil
}

// Play plays the currently targeted track.
func (i *ITunes) Play() error {
	if _, err := i.obj.CallMethod("Play"); err != nil {
		return err
	}
	return nil
}

// PreviousTrack returns to the previous track in the current playlist.
func (i *ITunes) PreviousTrack() error {
	if _, err := i.obj.CallMethod("PreviousTrack"); err != nil {
		return err
	}
	return nil
}

// QuitApp makes iTunes appilcation exit.
func (i *ITunes) QuitApp() error {
	if _, err := i.obj.CallMethod("Quit"); err != nil {
		return err
	}
	return nil
}

// Stop stops the playback.
func (i *ITunes) Stop() error {
	if _, err := i.obj.CallMethod("Stop"); err != nil {
		return err
	}
	return nil
}

// CurrentEncoder returns the currently selected encoder (AAC, MP3, AIFF, WAV, etc).
func (i *ITunes) CurrentEncoder() (string, error) {
	r, _ := i.obj.GetProperty("CurrentEncoder")
	enc := r.ToIDispatch()
	r, err := enc.GetProperty("Format")
	if err != nil {
		return "", err
	}
	return r.Value().(string), err
}

// Mute sets sound output mute state, depending on val (true sets mute on, false
// sets it off).
func (i *ITunes) Mute(val bool) error {
	if _, err := i.obj.PutProperty("Mute", val); err != nil {
		return err
	}
	return nil
}

// SoundVolume sets the sound output volue. You can set from 0 (minimum) to 100
// (maximum).
func (i *ITunes) SoundVolume(val int) error {
	ret, _ := i.obj.GetProperty("SoundVolumeControlEnabled")
	if ret.Value() != true {
		return fmt.Errorf("SoundVolume is disabled")
	}
	if val < 0 || val > 100 {
		return fmt.Errorf("Wrong value: 0 to 100 is acceptable")
	}
	if _, err := i.obj.PutProperty("SoundVolume", val); err != nil {
		return err
	}
	return nil
}

// SubscribeToPodcast subscribes to the specified podcast url u.
func (i *ITunes) SubscribeToPodcast(u string) error {
	escaped := url.PathEscape(u)
	if _, err := i.obj.CallMethod("SubscribeToPodcast", escaped); err != nil {
		return nil
	}
	return nil
}

// UpdatePodcastFeeds updates all podcast feeds.
func (i *ITunes) UpdatePodcastFeeds() error {
	if _, err := i.obj.CallMethod("UpdatePodcastFeeds"); err != nil {
		return nil
	}
	return nil
}

func (i *ITunes) ObjectPersistentID(obj interface{}) (int64, error) {
	var dispatch *ole.IDispatch
	if track, ok := obj.(*Track); ok {
		dispatch = track.obj
	} else if playlist, ok := obj.(*Playlist); ok {
		dispatch = playlist.obj
	} else {
		panic("object type is not supported")
	}

	var low, high int
	_, err := i.obj.CallMethod("GetITObjectPersistentIDs", dispatch, &high, &low)
	if err != nil {
		return 0, err
	}
	return int64(high)<<32 | int64(low), nil
}
