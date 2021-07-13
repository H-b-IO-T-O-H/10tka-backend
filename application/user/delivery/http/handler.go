package http

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	UserUseCase    user.UseCase
	SessionBuilder common.SessionBuilder
}

type reqUser struct {
	Id           string `json:"id"`
	Role         string    `json:"role" binding:"required"`
	Password     string    `json:"password"`
	Name         string    `json:"name" binding:"required"`
	Surname      string    `json:"surname" binding:"required"`
	Patronymic   string    `json:"patronymic"`
	Phone        string    `json:"phone" binding:"required" valid:"phone"`
	Email        string    `json:"email" binding:"required" valid:"email"`
	StudentGroup string    `json:"student_group"`
	Disciplines  []string  `json:"prof_disciplines"`
	About        string    `json:"about"`
}

type Resp struct {
	User *models.User `json:"user"`
}

type RespList struct {
	Users []models.User `json:"users"`
}

func NewRest(router *gin.RouterGroup, useCase user.UseCase, sessionBuilder common.SessionBuilder, AuthRequired gin.HandlerFunc, isAdminOrMethodist gin.HandlerFunc) *UserHandler {
	rest := &UserHandler{UserUseCase: useCase, SessionBuilder: sessionBuilder}
	rest.routes(router, AuthRequired, isAdminOrMethodist)
	return rest
}

func (u *UserHandler) routes(router *gin.RouterGroup, AuthRequired gin.HandlerFunc, isAdminOrMethodist gin.HandlerFunc) {
	router.POST("/login", u.LoginHandler)
	router.POST("/", u.CreateUserHandler)
	router.PUT("/", u.UpdateUserHandler)
	router.Use(AuthRequired)
	{
		router.GET("/me", u.GetCurrentUser)
		router.POST("/logout", u.LogoutHandler)
		router.GET("/students", u.GetStudents)
		router.GET("/professors", u.GetProfessors)
		//router.Use(isAdminOrMethodist)
		//{
		//
		//	router.GET("/students", u.GetStudents)
		//	router.GET("/professors", u.GetProfessors)
		//}
	}
}

func (u *UserHandler) LoginHandler(ctx *gin.Context) {
	var newUser models.UserLogin
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusForbidden, common.RespErr{Message: common.EmptyFieldErr})
		return
	}
	if err := common.ReqValidation(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, common.RespErr{Message: err.Error()})
		return
	}
	u.Login(ctx, newUser)
}

// Login
// @Summary Login
// @Description get user by username and password and returns userinfo with cookies
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} Resp
// @Failure 404 {object} common.Err
// @Failure 500 {object} common.Err
// @Router /users/login [post]
func (u *UserHandler) Login(ctx *gin.Context, newUser models.UserLogin) {
	buf, err := u.UserUseCase.Login(newUser)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}
	session := u.SessionBuilder.Build(ctx)
	if !newUser.ChekBox {
		session.Options(sessions.Options{Domain: "10-tka.ru",
			MaxAge:   2 * 24 * 3600,
			Secure:   true, // false for postman
			HttpOnly: true,
			Path:     "/",
			//SameSite: http.SameSiteNoneMode
			SameSite: http.SameSiteStrictMode,
		})
	}
	session.Set(common.UserRole, buf.Role)
	session.Set(common.UserId, buf.ID.String())
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.RespErr{Message: common.SessionErr})
		return
	}

	ctx.JSON(http.StatusOK, Resp{User: buf})
}

// LogoutHandler Logout
// @Summary Logout
// @Description clear user's session
// @Accept  json
// @Produce  json
// @Success 200 {object} nil
// @Failure 500 {object} common.Err
// @Router /users/logout [post]
func (u *UserHandler) LogoutHandler(ctx *gin.Context) {
	session := u.SessionBuilder.Build(ctx)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	err := session.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.RespErr{Message: common.SessionErr})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (u *UserHandler) CreateUserHandler(ctx *gin.Context) {
	var newUser reqUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, common.RespErr{Message: common.EmptyFieldErr})
		return
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err := u.UserUseCase.CreateUser(models.User{
		Role:         newUser.Role,
		Name:         newUser.Name,
		Surname:      newUser.Surname,
		Patronymic:   newUser.Patronymic,
		Email:        newUser.Email,
		Phone:        newUser.Phone,
		PasswordHash: passwordHash,
		StudentGroup: newUser.StudentGroup,
		Disciplines:  newUser.Disciplines,
		About:        newUser.About,
	}); err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (u *UserHandler) UpdateUserHandler(ctx *gin.Context) {
	var newUser reqUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, common.RespErr{Message: common.EmptyFieldErr})
		return
	}
	var passwordHash []byte

	if newUser.Password != "" {
		passwordHash, _ = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	}
	id, _ := uuid.Parse(newUser.Id)

	if err := u.UserUseCase.UpdateUser(models.User{
		ID:           id,
		Role:         newUser.Role,
		Name:         newUser.Name,
		Surname:      newUser.Surname,
		Patronymic:   newUser.Patronymic,
		Email:        newUser.Email,
		Phone:        newUser.Phone,
		PasswordHash: passwordHash,
		StudentGroup: newUser.StudentGroup,
		Disciplines:  newUser.Disciplines,
		About:        newUser.About,
	}); err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (u *UserHandler) GetCurrentUser(ctx *gin.Context) {
	session := u.SessionBuilder.Build(ctx)
	userID := session.Get(common.UserId)

	id, _ := uuid.Parse(userID.(string))
	userById, err := u.UserUseCase.GetUserById(id)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, Resp{User: userById})
}

func (u *UserHandler) GetStudents(ctx *gin.Context) {
	users, err := u.UserUseCase.GetUsersAll(common.Student)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, RespList{Users: users})
}

func (u *UserHandler) GetProfessors(ctx *gin.Context) {
	users, err := u.UserUseCase.GetUsersAll(common.Professor)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, RespList{Users: users})
}
