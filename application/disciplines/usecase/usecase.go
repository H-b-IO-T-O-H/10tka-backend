package usecase

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/disciplines"
)

type DisciplineUseCase struct {
	repos  disciplines.RepositoryDiscipline
}

func NewDisciplineUseCase(repos disciplines.RepositoryDiscipline) disciplines.UseCase {
	return &DisciplineUseCase{repos:  repos}
}

func (d DisciplineUseCase) GetDisciplines() ([]models.Discipline, common.Err) {
	return d.repos.GetDisciplines()
}

