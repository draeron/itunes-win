package itunes

func (p *PlaylistCollection) Count() (int64, error) {
	c, err := p.getInt64Property("Count")
	return c, err
}

func (p *PlaylistCollection) ByIndex(idx int) (*Playlist, error) {
	obj, err := p.getObjectProperty("Item", idx)
	if err != nil || obj == nil || obj.IsNil() {
		return nil, err
	}
	return &Playlist{*obj}, nil
}

func (p *PlaylistCollection) ByName(name string) (*Playlist, error) {
	obj, err := p.getObjectProperty("ItemByName", name)
	if err != nil || obj == nil || obj.IsNil() {
		return nil, err
	}
	return &Playlist{*obj}, nil
}

func (p *PlaylistCollection) ByPersistentID(pid PersistentID) (*Playlist, error) {
	obj, err := p.getObjectByPersistentID(pid)
	if err != nil || obj == nil || obj.IsNil() {
		return nil, err
	}
	return &Playlist{*obj}, nil
}
