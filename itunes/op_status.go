package itunes

type OperationStatus struct {
  COM
}

func (o *OperationStatus) GetTracks() (*TrackCollection, error) {
  ret, err := o.COM.getObjectProperty("Tracks")
  if err != nil {
    return nil, err
  }
  return &TrackCollection{*ret}, nil
}