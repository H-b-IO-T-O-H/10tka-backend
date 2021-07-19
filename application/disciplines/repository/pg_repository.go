package repository

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/disciplines"
	"gorm.io/gorm"
	"net/http"
)

type pgStorage struct {
	db *gorm.DB
}

func NewPgRepository(db *gorm.DB) disciplines.RepositoryDiscipline {
	return &pgStorage{db: db}
}

func (p pgStorage) GetDisciplines() ([]models.Discipline, common.Err) {
	var listOfDisciplines []models.Discipline

	err := p.db.Raw("select * from public.disciplines").Scan(&listOfDisciplines).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RespErr{Message: common.NotFound, Status: http.StatusNotFound}
		}
		return nil, common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	return listOfDisciplines, nil
}
