package http

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/disciplines"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DisciplineHandler struct {
	DisciplineUseCase disciplines.UseCase
	SessionBuilder    common.SessionBuilder
}

type Resp struct {
	Discipline *models.Discipline `json:"discipline"`
}

type RespList struct {
	Disciplines []models.Discipline `json:"disciplines"`
}

func NewRest(router *gin.RouterGroup, useCase disciplines.UseCase, sessionBuilder common.SessionBuilder, AuthRequired gin.HandlerFunc) *DisciplineHandler {
	rest := &DisciplineHandler{DisciplineUseCase: useCase, SessionBuilder: sessionBuilder}
	rest.routes(router, AuthRequired)
	return rest
}

func (d *DisciplineHandler) routes(router *gin.RouterGroup, AuthRequired gin.HandlerFunc) {
	router.GET("/all", d.GetAllDisciplinesHandler)

}

func (d *DisciplineHandler) GetAllDisciplinesHandler(ctx *gin.Context) {
	list, err := d.DisciplineUseCase.GetDisciplines()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}
	ctx.JSON(http.StatusOK, RespList{Disciplines: list})
}
