package itunes

// GetCount returns the number of track in this tracks.
func (t *Tracks) GetCount() (int64, error) {
	v, err := t.obj.GetProperty("Count")
	if err != nil {
		return 0, err
	}
	return v.Val, nil
}

// GetTrackByIndex returns Track in t, corresponding to the given index (1-based).
func (t *Tracks) GetTrackByIndex(index int) (*Track, error) {
	r, err := t.obj.GetProperty("Item", index)
	if err != nil {
		return nil, err
	}
	o := COM{
		obj: r.ToIDispatch(),
	}
	return &Track{COM: o}, nil
}

// GetTrackByName returns Track in t with the specified name.
func (t *Tracks) GetTrackByName(name string) (*Track, error) {
	r, err := t.obj.GetProperty("ItemByName", name)
	if err != nil {
		return nil, err
	}
	o := COM{
		obj: r.ToIDispatch(),
	}
	return &Track{COM: o}, nil
}
