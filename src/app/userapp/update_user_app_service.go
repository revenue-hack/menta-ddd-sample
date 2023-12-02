package userapp

import (
	"context"

	"github.com/revenue-hack/menta-ddd-sample/src/domain/userdm"
)

type updateUserAppService struct {
	userRepo userdm.UserRepository
}

type UpdateUserRequest struct {
	ID        string
	FirstName string
	LastName  string
	Skills    []UpdateSkillRequest
	Careers   []UpdateCareerRequest
}

type UpdateSkillRequest struct {
	ID    *string
	TagID string
}
type UpdateCareerRequest struct {
	ID     *string
	Detail string
}

func NewUpdateUserAppService(userRepo userdm.UserRepository) *updateUserAppService {
	return &updateUserAppService{userRepo}
}

func (app *updateUserAppService) Exec(ctx context.Context, req *UpdateUserRequest) error {
	userID, err := userdm.NewUserIDByVal(req.ID)
	if err != nil {
		return err
	}
	user, err := app.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	skills := make([]userdm.SkillParamIfUpdate, len(req.Skills))
	careers := make([]userdm.CareerParamIfUpdate, len(req.Careers))

	for i, reqCareer := range req.Careers {
		careers[i] = userdm.CareerParamIfUpdate{
			ID:     reqCareer.ID,
			Detail: reqCareer.Detail,
		}
	}

	for i, reqSkill := range req.Skills {
		skills[i] = userdm.SkillParamIfUpdate{
			ID:    reqSkill.ID,
			TagID: reqSkill.TagID,
		}
	}

	if err := user.Update(
		req.FirstName,
		req.LastName,
		careers,
		skills,
	); err != nil {
		return err
	}

	return app.userRepo.Update(ctx, user)
}
