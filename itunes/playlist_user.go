package itunes

import (
	"fmt"
	"time"
)

func (p *Playlist) isUserList() bool {
	return p.Kind() == ITPlaylistKindUser
}

func (p *Playlist) IsSmart() (bool, error) {
	if !p.isUserList() {
		return false, fmt.Errorf("Not a User playlist")
	}
	return p.getBoolProperty("Smart")
}

func (p *Playlist) CreateFolder(name string) (*Playlist, error) {
	if !p.isUserList() {
		return nil, fmt.Errorf("Not a User playlist")
	}

	ret, err := p.obj.CallMethod("CreateFolder", name)
	if err != nil {
		return nil, err
	}
	return &Playlist{COM{ret.ToIDispatch()}}, nil
}

func (p *Playlist) CreatePlaylist(name string) (*Playlist, error) {
	if !p.isUserList() {
		return nil, fmt.Errorf("Not a User playlist")
	}

	ret, err := p.obj.CallMethod("CreatePlaylist", name)
	if err != nil {
		return nil, err
	}
	return &Playlist{COM{ret.ToIDispatch()}}, nil
}

func (p *Playlist) Parent() (*Playlist, error) {
	if !p.isUserList() {
		return nil, fmt.Errorf("Not a User playlist")
	}

	obj, err := p.getObjectProperty("Parent")
	if err != nil {
		return nil, err
	}
	if obj != nil && obj.IsNil() {
		return nil, nil
	}
	return &Playlist{*obj}, nil
}

func (p *Playlist) Size() (uint64, error) {
	ret, err := p.getDoubleProperty("Size")
	return uint64(ret), err
}

func (p *Playlist) Shuffle() (bool, error) {
	return p.getBoolProperty("Shuffle")
}

func (p *Playlist) Duration() (time.Duration, error) {
	return p.getDurationProperty("Duration")
}

func (p *Playlist) SetParent(parent *Playlist) error {
	if !p.isUserList() {
		return fmt.Errorf("Not a User playlist")
	}

	_, err := p.obj.PutProperty("Parent", parent.obj)
	return err
}

func (p *Playlist) AddTrack(track *Track) error {
	if !p.isUserList() {
		return fmt.Errorf("Not a User playlist")
	}

	_, err := p.obj.CallMethod("AddTrack", track.COM.obj)
	return err
}

func (p *Playlist) SpecialKind() (ITUserPlaylistSpecialKind, error) {
	if !p.isUserList() {
		return ITUserPlaylistSpecialKindNone, fmt.Errorf("Not a User playlist")
	}

	ret, err := p.getInt64Property("SpecialKind")
	if err != nil {
		return ITUserPlaylistSpecialKindNone, err
	}
	return ITUserPlaylistSpecialKind(ret), nil
}
