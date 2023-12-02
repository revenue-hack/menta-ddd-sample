package userdm

import (
	"time"

	"github.com/revenue-hack/menta-ddd-sample/src/domain/shared"
	"github.com/revenue-hack/menta-ddd-sample/src/domain/tagdm"
)

type CareerParamIfCreate struct {
	Detail string
}
type SkillParamIfCreate struct {
	TagID string
}

func GenIfCreate(
	first,
	last string,
	reqCareers []CareerParamIfCreate,
	reqSkills []SkillParamIfCreate,
) (*User, error) {
	careers := make([]Career, len(reqCareers))
	for i, rc := range reqCareers {
		c, err := newCareer(NewCareerID(), rc.Detail)
		if err != nil {
			return nil, err
		}
		careers[i] = *c
	}
	skills := make([]Skill, len(reqSkills))
	for i, rs := range reqSkills {
		tagID, err := tagdm.NewTagIDByVal(rs.TagID)
		if err != nil {
			return nil, err
		}
		s, err := newSkill(NewSkillID(), tagID)
		if err != nil {
			return nil, err
		}
		skills[i] = *s
	}

	return newUser(NewUserID(), first, last, shared.NewCreatedAt(), time.Now(), careers, skills)
}
func GenForTest(
	id UserID,
	first,
	last string,
	careers []Career,
	skills []Skill,
) (*User, error) {
	return newUser(id, first, last, shared.NewCreatedAt(), time.Now(), careers, skills)
}
