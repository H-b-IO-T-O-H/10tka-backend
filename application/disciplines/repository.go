package disciplines

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
)

type RepositoryDiscipline interface {
	GetDisciplines() ([]models.Discipline, common.Err)
}
