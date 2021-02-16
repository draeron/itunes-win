package itunes

import (
	"strconv"
)

type PersistentID uint64

func (p PersistentID) String() string {
	return strconv.FormatUint(uint64(p), 16)
}

func (p PersistentID) High() int {
	return int(p >> 32)
}

func (p PersistentID) Low() int {
	return int(p)
}

func ParsePersistentID(str string) (PersistentID, error) {
	id, err := strconv.ParseUint(str[:16], 16, 64)
	if err != nil {
		return 0, err
	}
	return PersistentID(id), err
}
