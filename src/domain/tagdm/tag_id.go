package tagdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type TagID string

func NewTagID() TagID {
	return TagID(uuid.New().String())
}
func NewTagIDByVal(val string) (TagID, error) {
	if val == "" {
		return TagID(""), xerrors.New("tag id must not be empty")
	}
	return TagID(val), nil
}

func (id TagID) String() string {
	return string(id)
}

func (id TagID) Equal(TagID2 TagID) bool {
	return string(id) == string(TagID2)
}
