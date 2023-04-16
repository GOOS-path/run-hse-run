package conversions

import (
	"Run_Hse_Run/genproto"
	"Run_Hse_Run/pkg/model"
)

func ConvertRoom(room model.Room) *genproto.Room {
	return &genproto.Room{
		Id:       room.Id,
		Code:     room.Code,
		CampusId: room.CampusId,
	}
}
