package repository

import (
	"fmt"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/smtp"
)

//; alter sequence public.requests_req_id_seq restart with 1
type pgStorage struct {
	db *gorm.DB
}

func NewPgRepository(db *gorm.DB) user.RepositoryUser {
	return &pgStorage{db: db}
}

func (p *pgStorage) Login(user models.UserLogin) (*models.User, common.Err) {
	//userDB := new(models.User)
	//
	//err := p.db.Take(userDB, "email = ?", user.Email).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return nil, common.RespErr{Message: common.AuthErr, Status: http.StatusNotFound}
	//	}
	//	return nil, common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//}
	//// compare password with the hashed one
	//err = bcrypt.CompareHashAndPassword(userDB.PasswordHash, []byte(user.Password))
	//if err != nil {
	//	return nil, common.RespErr{Message: common.AuthErr, Status: http.StatusNotFound}
	//}
	//if userDB.Role == common.Student {
	//	err = p.db.Raw("select g.group_name from public.groups g join students s on s.group_id=g.group_id where s.user_id=?", userDB.ID).Row().Scan(&userDB.StudentGroup)
	//	if err != nil {
	//		return nil, common.RespErr{Message: common.AuthErr, Status: http.StatusNotFound}
	//	}
	//}
	//return userDB, nil
	return nil, nil
}

func (p *pgStorage) CreateUser(newUser models.User) common.Err {
	//tx := p.db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//if err := tx.Error; err != nil {
	//	return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//}
	//
	//newUser.DisciplinesOrm = "{}"
	//if newUser.About == "" {
	//	newUser.About = "-"
	//}
	//if newUser.Role == common.Professor {
	//	newUser.DisciplinesOrm = "{"
	//	for idx, dis := range newUser.Disciplines {
	//		query := fmt.Sprintf("insert into public.disciplines(discipline, prof_cnt) values ('%s', '%d') on conflict(discipline) do update set prof_cnt=disciplines.prof_cnt+1", dis, 1)
	//		err := tx.Exec(query).Error
	//		if err != nil {
	//			return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//		}
	//		if idx != len(newUser.Disciplines)-1 {
	//			newUser.DisciplinesOrm += fmt.Sprintf("%s, ", dis)
	//		} else {
	//			newUser.DisciplinesOrm += fmt.Sprintf("%s", dis)
	//		}
	//	}
	//	newUser.DisciplinesOrm += "}"
	//}
	//if err := tx.Create(&newUser).Error; err != nil {
	//	msg := err.Error()
	//	if common.RecordExists(msg) {
	//		return common.RespErr{Message: common.UserExistErr, Status: http.StatusConflict}
	//	}
	//	return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//}
	//
	//if newUser.Role == common.Student {
	//	group := new(models.Group)
	//	group.GroupName = newUser.StudentGroup
	//	err := tx.Take(&group).Where("group_name = ?", group.GroupName).Error
	//	if err != nil {
	//		if err == gorm.ErrRecordNotFound {
	//			if tx.Create(group).Error != nil {
	//				return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	//			}
	//		} else {
	//			return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	//		}
	//	}
	//	err = tx.Exec(fmt.Sprintf("insert into public.students(group_id, user_id) values ('%s', '%s')", group.GroupId, newUser.ID)).Error
	//	if err != nil {
	//		return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//	}
	//}
	//if err := tx.Commit().Error; err != nil {
	//	return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//}
	return nil
}

func (p *pgStorage) GetUserById(userId uuid.UUID) (*models.User, common.Err) {
	//userDB := new(models.User)
	//
	//userDB.ID = userId
	//if err := p.db.Take(userDB).Error; err != nil {
	//	msg := err.Error()
	//	if common.NoRows(msg) {
	//		return nil, common.RespErr{Message: common.AuthErr, Status: http.StatusNotFound}
	//	} else {
	//		return nil, common.RespErr{Status: http.StatusInternalServerError, Message: msg}
	//	}
	//}
	//return userDB, nil
	return nil, nil
}


func test() {
	auth := smtp.PlainAuth("", "vladislav.amelin@2015@yandex", "extremely_secret_pass", "smtp.yandex.ru")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"importantv@yandex.ru"}
	msg := []byte("To: bill@gates.com\r\n" +
		"Subject: Why are you not using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here’s the space for our great sales pitch\r\n")
	err := smtp.SendMail("smtp.yandex.ru:465", auth, "vladislav.amelin@2015@yandex.ru", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}



func (p *pgStorage) UpdateUser(newUser models.User) common.Err {
	//var passwd string
	//
	////test()
	//tx := p.db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//if err := tx.Error; err != nil {
	//	return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//}
	//err := tx.Raw(fmt.Sprintf("select password_hash from public.users where user_id='%s'", newUser.ID)).Scan(&passwd).Error
	//newUser.PasswordHash = []byte(passwd)
	//if err != nil{
	//	return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	//}
	//err = tx.Exec(fmt.Sprintf("delete from public.users where user_id = '%s'", newUser.ID)).Error
	//if err != nil {
	//	return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	//}
	//newUser.DisciplinesOrm = "{}"
	//if newUser.About == "" {
	//	newUser.About = "-"
	//}
	//if newUser.Role == common.Professor {
	//	newUser.DisciplinesOrm = "{"
	//	for idx, dis := range newUser.Disciplines {
	//		query := fmt.Sprintf("insert into public.disciplines(discipline, prof_cnt) values ('%s', '%d') on conflict(discipline) do update set prof_cnt=disciplines.prof_cnt+1", dis, 1)
	//		err := tx.Exec(query).Error
	//		if err != nil {
	//			return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//		}
	//		if idx != len(newUser.Disciplines)-1 {
	//			newUser.DisciplinesOrm += fmt.Sprintf("%s, ", dis)
	//		} else {
	//			newUser.DisciplinesOrm += fmt.Sprintf("%s", dis)
	//		}
	//	}
	//	newUser.DisciplinesOrm += "}"
	//}
	//if err := tx.Create(&newUser).Error; err != nil {
	//	msg := err.Error()
	//	if common.RecordExists(msg) {
	//		return common.RespErr{Message: common.UserExistErr, Status: http.StatusConflict}
	//	}
	//	return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//}
	//
	//if newUser.Role == common.Student {
	//	group := new(models.Group)
	//	group.GroupName = newUser.StudentGroup
	//	err := tx.Take(&group).Where("group_name = ?", group.GroupName).Error
	//	if err != nil {
	//		if err == gorm.ErrRecordNotFound {
	//			if tx.Create(group).Error != nil {
	//				return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	//			}
	//		} else {
	//			return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	//		}
	//	}
	//	err = tx.Exec(fmt.Sprintf("insert into public.students(group_id, user_id) values ('%s', '%s')", group.GroupId, newUser.ID)).Error
	//	if err != nil {
	//		return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//	}
	//}
	//if err := tx.Commit().Error; err != nil {
	//	return common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	//}
	return nil
}

func (p *pgStorage) GetUsersAll(userType string) ([]models.User, common.Err) {
	var usersList []models.User

	err := p.db.Raw(fmt.Sprintf("select * from public.users where role = '%s'", userType)).Scan(&usersList).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RespErr{Message: common.NotFound, Status: http.StatusNotFound}
		}
		return nil, common.RespErr{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	return usersList, nil
}

func (p *pgStorage) GetUsers(userType string, start uint8, limit uint8) ([]models.User, common.Err) {
	panic("implement me")
}
