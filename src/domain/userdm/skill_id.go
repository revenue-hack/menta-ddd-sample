package userdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type SkillID string

func NewSkillID() SkillID {
	return SkillID(uuid.New().String())
}
func NewSkillIDByVal(val string) (SkillID, error) {
	if val == "" {
		return SkillID(""), xerrors.New("skill id must not be empty")
	}
	return SkillID(val), nil
}

func (id SkillID) String() string {
	return string(id)
}

func (id SkillID) Equal(SkillID2 SkillID) bool {
	return string(id) == string(SkillID2)
}
