package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type User struct {
	ID           int    `gorm:"column:user_id" json:"id"`
	Role         string `gorm:"column:role" json:"role"`
	Email        string `gorm:"column:email" json:"email"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Name         string `gorm:"column:name" json:"name"`
	Surname      string `gorm:"column:surname" json:"surname"`
	Patronymic   string `gorm:"column:patronymic" json:"patronymic"`
	PasswordHash []byte `gorm:"column:password_hash" json:"-"`
	BirthDate    string `gorm:"column:birth_date" json:"birth_date"`
	About        string `gorm:"column:about" json:"about"`
}

type Students struct {
	ID            int         `gorm:"column:user_id" json:"id"`
	OrgCuratorId  int         `gorm:"column:org_curator_id" json:"org_curator_id"`
	GroupNmb      pgtype.Int2 `gorm:"column:group_nmb" json:"group_nmb"`
	AdmissionDate string      `gorm:"column:admission_date" json:"admission_date"`
	IsGraduated   bool        `gorm:"column:is_graduated" json:"is_graduated"`
	InAcadem      bool        `gorm:"column:in_academ" json:"in_academ"`
}

type Groups struct {
	GroupNmb       pgtype.Int2 `gorm:"column:group_nmb" json:"group_nmb"`
	GroupElderId   int         `gorm:"column:group_elder_id" json:"group_elder_id"`
	TimetableId    uuid.UUID   `gorm:"column:timetable_id" json:"timetable_id"`
	GroupCuratorId int         `gorm:"column:group_curator_id" json:"group_curator_id"`
	Semester       pgtype.Int2 `gorm:"column:semester" json:"semester"`
	StudentsCnt    pgtype.Int2 `gorm:"column:students_cnt" json:"students_cnt"`
}

type Professors struct {
	ID             int          `gorm:"column:user_id" json:"user_id"`
	Seniority      pgtype.Int2  `gorm:"column:seniority" json:"seniority"`
	AcademicDegree string       `gorm:"column:academic_degree" json:"academic_degree"`
	Rank           string       `gorm:"column:rank" json:"rank"`
	ContestDate    string       `gorm:"column:contest_date" json:"contest_date"`
	IsCombining    bool         `gorm:"column:is_combining" json:"is_combining"`
	SharedHours    pgtype.Int2  `gorm:"column:shared_hours" json:"shared_hours"`
	WorkRate       pgtype.Int4  `gorm:"column:work_rate" json:"work_rate"`
	WorkTime       string       `gorm:"column:work_time" json:"work_time"`
	Disciplines    []Discipline `gorm:"column:disciplines" json:"disciplines"`
}
type Competenties struct {
	Competention string `gorm:"column:competention" json:"competention"`
	UsersIds     []User `gorm:"column:users_ids" json:"users_ids"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ChekBox  bool   `json:"checkbox"`
}

func (u User) TableName() string {
	return "public.users"
}

//easyjson:json
type UsersList []User
