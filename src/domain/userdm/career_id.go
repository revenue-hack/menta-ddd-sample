package userdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type CareerID string

func NewCareerID() CareerID {
	return CareerID(uuid.New().String())
}
func NewCareerIDByVal(val string) (CareerID, error) {
	if val == "" {
		return CareerID(""), xerrors.New("career id must not be empty")
	}
	return CareerID(val), nil
}

func (id CareerID) String() string {
	return string(id)
}

func (id CareerID) Equal(CareerID2 CareerID) bool {
	return string(id) == string(CareerID2)
}
