package models

type Discipline struct {
	DisciplineId      int16  `gorm:"column:discipline_id" json:"discipline_id"`
	DisciplineContent string `gorm:"column:discipline" json:"discipline_content"`
	ProfCnt           string `gorm:"column:prof_cnt" json:"prof_cnt"`
}

func (d Discipline) TableName() string {
	return "public.disciplines"
}
