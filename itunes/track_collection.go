package itunes

// GetCount returns the number of track in this tracks.
func (t *TrackCollection) Count() (int64, error) {
	v, err := t.obj.GetProperty("Count")
	if err != nil {
		return 0, err
	}
	return v.Val, nil
}

// GetTrackByIndex returns Track in t, corresponding to the given index (1-based).
func (t *TrackCollection) ByIndex(index int) (*Track, error) {
	obj, err := t.getObjectProperty("Item", index)
	if err != nil || obj == nil || obj.IsNil() {
		return nil, err
	}
	return &Track{*obj}, nil
}

// GetTrackByName returns Track in t with the specified name.
func (t *TrackCollection) ByName(name PersistentID) (*Track, error) {
	obj, err := t.getObjectProperty("ItemByName", name)
	if err != nil || obj == nil || obj.IsNil() {
		return nil, err
	}
	return &Track{*obj}, nil
}

// GetTrackByName returns Track in t with the specified name.
func (t *TrackCollection) ByPersistentID(pid PersistentID) (*Track, error) {
	obj, err := t.getObjectByPersistentID(pid)
	if err != nil || obj == nil || obj.IsNil() {
		return nil, err
	}
	return &Track{*obj}, nil
}
