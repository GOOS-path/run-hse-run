package conversions

import (
	"Run_Hse_Run/genproto"
	"Run_Hse_Run/pkg/model"
)

func ConvertUser(user model.User) *genproto.User {
	return &genproto.User{
		Id:       user.Id,
		Nickname: user.Nickname,
		Email:    user.Email,
		Image:    user.Image,
		Score:    user.Score,
	}
}
