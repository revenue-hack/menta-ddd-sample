package userapp

import (
	"context"

	"github.com/revenue-hack/menta-ddd-sample/src/domain/userdm"
)

type createUserAppService struct {
	userRepo userdm.UserRepository
}

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Skills    []CreateSkillRequest
	Careers   []CreateCareerRequest
}

type CreateSkillRequest struct {
	TagID string
}
type CreateCareerRequest struct {
	Detail string
}

func NewCreateUserAppService(userRepo userdm.UserRepository) *createUserAppService {
	return &createUserAppService{userRepo}
}

func (app *createUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	careers := make([]userdm.CareerParamIfCreate, len(req.Careers))
	skills := make([]userdm.SkillParamIfCreate, len(req.Skills))

	for i, reqCareer := range req.Careers {
		careers[i] = userdm.CareerParamIfCreate{
			Detail: reqCareer.Detail,
		}
	}

	for i, reqSkill := range req.Skills {
		skills[i] = userdm.SkillParamIfCreate{
			TagID: reqSkill.TagID,
		}
	}

	user, err := userdm.GenIfCreate(req.FirstName, req.LastName, careers, skills)
	if err != nil {
		return err
	}

	return app.userRepo.Store(ctx, user)
}
