package usecase

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/user"
	"github.com/apsdehal/go-logger"
	"github.com/google/uuid"
)

type UserUseCase struct {
	iLog   *logger.Logger
	errLog *logger.Logger
	repos  user.RepositoryUser
}

func NewUserUseCase(iLog *logger.Logger, errLog *logger.Logger,
	repos user.RepositoryUser) *UserUseCase {
	return &UserUseCase{
		iLog:   iLog,
		errLog: errLog,
		repos:  repos,
	}
}

func (u *UserUseCase) Login(user models.UserLogin) (*models.User, common.Err) {
	return u.repos.Login(user)
}

func (u *UserUseCase) CreateUser(newUser models.User) common.Err {
	return u.repos.CreateUser(newUser)
}

func (u *UserUseCase) UpdateUser(newUser models.User) common.Err {
	return u.repos.UpdateUser(newUser)
}

func (u *UserUseCase) GetUserById(userId uuid.UUID) (*models.User, common.Err) {
	return u.repos.GetUserById(userId)
}

func (u *UserUseCase) GetUsers(userType string, start uint8, limit uint8) ([]models.User, common.Err) {
	return u.repos.GetUsers(userType, start, limit)
}

func (u *UserUseCase) GetUsersAll(userType string) ([]models.User, common.Err) {
	return u.repos.GetUsersAll(userType)
}