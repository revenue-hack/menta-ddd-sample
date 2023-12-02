package userdm

import "github.com/revenue-hack/menta-ddd-sample/src/domain/tagdm"

type Skill struct {
	id    SkillID
	tagID tagdm.TagID
}

func newSkill(id SkillID, tagID tagdm.TagID) (*Skill, error) {
	// validationチェックが入る
	return &Skill{
		id:    id,
		tagID: tagID,
	}, nil
}
