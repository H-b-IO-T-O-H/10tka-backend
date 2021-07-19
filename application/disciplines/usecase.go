package disciplines

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
)

type UseCase interface {
	GetDisciplines() ([]models.Discipline, common.Err)
}
