package userdm

import (
	"time"
	"unicode/utf8"

	"github.com/revenue-hack/menta-ddd-sample/src/domain/shared"
	"github.com/revenue-hack/menta-ddd-sample/src/domain/tagdm"
	"golang.org/x/xerrors"
)

var (
	validationSkillLen = 0
)

type User struct {
	id        UserID
	firstName string
	lastName  string
	skills    []Skill
	careers   []Career
	createdAt shared.CreatedAt
	updatedAt time.Time
}

func (u *User) ID() UserID {
	return u.id
}
func (u *User) FirstName() string {
	return u.firstName
}
func (u *User) LastName() string {
	return u.lastName
}
func (u *User) CreatedAt() shared.CreatedAt {
	return u.createdAt
}

var (
	firstNameLength = 30
	lastNameLength  = 30
)

func newUser(
	id UserID,
	first,
	last string,
	createdAt shared.CreatedAt,
	updatedAt time.Time,
	careers []Career,
	skills []Skill,
) (*User, error) {
	if first == "" {
		return nil, xerrors.New("first name must be not empty")
	}
	if last == "" {
		return nil, xerrors.New("last name must be not empty")
	}

	if l := utf8.RuneCountInString(first); l > firstNameLength {
		return nil, xerrors.Errorf("first name must be less than %d", firstNameLength)
	}
	if l := utf8.RuneCountInString(last); l > lastNameLength {
		return nil, xerrors.Errorf("last name must be less than %d", lastNameLength)
	}

	// 最低一つなければ作成できない
	if len(skills) > validationSkillLen {
		return nil, xerrors.Errorf("skill must be more than one")
	}

	return &User{
		id:        id,
		firstName: first,
		lastName:  last,
		createdAt: createdAt,
		updatedAt: updatedAt,
		skills:    skills,
		careers:   careers,
	}, nil
}

type CareerParamIfUpdate struct {
	ID     *string
	Detail string
}
type SkillParamIfUpdate struct {
	ID    *string
	TagID string
}

func (u *User) Update(
	first,
	last string,
	reqCareers []CareerParamIfUpdate,
	reqSkills []SkillParamIfUpdate,
) error {
	careers := make([]Career, len(reqCareers))
	for i, rc := range reqCareers {
		var career *Career
		if rc.ID != nil {
			id, err := NewCareerIDByVal(*rc.ID)
			if err != nil {
				return err
			}
			c, err := newCareer(id, rc.Detail)
			career = c
		} else {
			c, err := newCareer(NewCareerID(), rc.Detail)
			if err != nil {
				return err
			}
			career = c
		}
		careers[i] = *career
	}
	skills := make([]Skill, len(reqSkills))
	for i, rs := range reqSkills {
		var skill *Skill
		if rs.ID != nil {
			id, err := NewSkillIDByVal(*rs.ID)
			if err != nil {
				return err
			}
			tagID, err := tagdm.NewTagIDByVal(rs.TagID)
			if err != nil {
				return err
			}
			s, err := newSkill(id, tagID)
			if err != nil {
				return err
			}
			skill = s
		} else {
			tagID, err := tagdm.NewTagIDByVal(rs.TagID)
			if err != nil {
				return err
			}
			s, err := newSkill(NewSkillID(), tagID)
			if err != nil {
				return err
			}
			skill = s
		}

		skills[i] = *skill
	}

	// 最低一つなければ作成できない
	if len(skills) > validationSkillLen {
		return xerrors.Errorf("skill must be more than one")
	}
	u.careers = careers
	u.skills = skills

	u.firstName = first
	u.lastName = last

	return nil
}
