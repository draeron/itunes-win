package itunes

import (
	"math"
	"time"
)

func (c *COM) Name() string {
	n, _ := c.getStringProperty("Name")
	return n
}

func (c *COM) getInt64Property(property string) (int64, error) {
	n, err := c.obj.GetProperty(property)
	if err != nil {
		return 0, err
	}
	return n.Val, nil
}

func (c *COM) getDoubleProperty(property string) (float64, error) {
	n, err := c.obj.GetProperty(property)
	if err != nil {
		return 0, err
	}

	// b := make([]byte, 8)
	// binary.LittleEndian.PutUint64(b, uint64(n.Val))
	float := math.Float64frombits(uint64(n.Val))
	return float, nil
}

func (c *COM) getStringProperty(property string) (string, error) {
	n, err := c.obj.GetProperty(property)
	if err != nil {
		return "", err
	}
	return n.Value().(string), nil
}

func (c *COM) getObjectProperty(property string, params ...interface{}) (*COM, error) {
	n, err := c.obj.GetProperty(property, params...)
	if err != nil {
		return nil, err
	}
	return &COM{n.ToIDispatch()}, err
}

func (c *COM) setStringProperty(property string, value string) error {
	_, err := c.obj.PutProperty(property, value)
	return err
}

func (c *COM) setInt64Property(property string, value string) error {
	_, err := c.obj.PutProperty(property, value)
	return err
}

func (c *COM) getDateProperty(property string) (time.Time, error) {
	n, err := c.obj.GetProperty(property)
	if err != nil {
		return time.Time{}, err
	}
	return n.Value().(time.Time), nil
}

func (c *COM) getDurationProperty(property string) (time.Duration, error) {
	seconds, err := c.getInt64Property("Duration")
	if err != nil {
		return 0, err
	}
	return time.Second * time.Duration(seconds), nil
}

func (c *COM) getBoolProperty(property string) (bool, error) {
	n, err := c.obj.GetProperty(property)
	if err != nil {
		return false, err
	}
	return n.Value().(bool), nil
}

func (c *COM) setProperty(property string, val interface{}) error {
	_, err := c.obj.PutProperty(property, val)
	if err != nil {
		return err
	}
	return nil
}

func (c *COM) getKind() int64 {
	val, _ := c.getInt64Property("Kind")
	return val
}

// GetTrackByName returns Track in t with the specified name.
func (c *COM) getObjectByPersistentID(pid PersistentID) (*COM, error) {
	return c.getObjectProperty("ItemByPersistentID", pid.High(), pid.Low())
}

func (c *COM) ObjectIDs() (sourceID int, playlistID int, trackID int, databaseID int, err error) {
	_, err = c.obj.CallMethod("GetITObjectIDs", &sourceID, &playlistID, &trackID, &databaseID)
	return
}
