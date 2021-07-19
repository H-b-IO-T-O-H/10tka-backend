package http

import (
	"fmt"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/posts"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path"
	"strings"
)

type PostHandler struct {
	PostUSeCase    posts.UseCase
	SessionBuilder common.SessionBuilder
}

type RespImg struct {
	Location string `json:"location"`
}

func NewRest(router *gin.RouterGroup, useCase posts.UseCase, sessionBuilder common.SessionBuilder, AuthRequired gin.HandlerFunc, isAdminOrMethodist gin.HandlerFunc) *PostHandler {
	rest := &PostHandler{PostUSeCase: useCase, SessionBuilder: sessionBuilder}
	rest.routes(router, AuthRequired, isAdminOrMethodist)
	return rest
}

func (p *PostHandler) routes(router *gin.RouterGroup, AuthRequired gin.HandlerFunc, isAdminOrMethodist gin.HandlerFunc) {
	router.GET("/current-id", p.GetCurrentPostId)
	router.POST("/:post-id/upload-image", p.UploadPostsImages)
	router.POST("/", p.CreatePost)
	router.GET("/:post-id/download-image/:image-name", p.DownloadPostImage)
	router.GET("/:post-id", p.GetPostById)
	router.Use(AuthRequired)
	{
	}
}

func (p *PostHandler) GetCurrentPostId(ctx *gin.Context) {
	id, err := p.PostUSeCase.GetCurrentPostId()
	if err != nil {
		ctx.JSON(err.StatusCode(), err.Msg())
	}

	ctx.JSON(http.StatusOK, id)
}

func (p *PostHandler) UploadPostsImages(ctx *gin.Context) {
	postId := strings.Split(ctx.Request.RequestURI, "/")[4]
	if err := ctx.Request.ParseMultipartForm(2 >> 20); err != nil {
		ctx.JSON(http.StatusBadRequest, common.EmptyFieldErr)
	}
	file, header, _ := ctx.Request.FormFile("file")

	if !common.IsValidFile(header, []string{common.JpegMime, common.PngMime}) {
		ctx.JSON(http.StatusBadRequest, common.EmptyFieldErr)
		return
	}
	extension := strings.Split(header.Header.Get("Content-type"), "/")[1]
	pathDir, err := common.CreateDir(fmt.Sprintf("post_%s_dir", postId))
	imgName := fmt.Sprintf("%s.%s", uuid.New().String(), extension)
	_, err = common.AddOrUpdateUserFile(file, fmt.Sprintf("%s/%s", pathDir, imgName))
	if err != nil {
		ctx.JSON(err.StatusCode(), err.Msg())
		return
	}
	location := fmt.Sprintf("%s/posts/%s/download-image/%s", common.Domain, postId, imgName)
	ctx.JSON(http.StatusOK, RespImg{Location: location})
}

func (p *PostHandler) DownloadPostImage(ctx *gin.Context) {
	postId := ctx.Param("post-id")
	imgName := ctx.Param("image-name")
	targetPath := path.Join(common.PathToSaveStatic, fmt.Sprintf("post_%s_dir", postId), imgName)
	if !common.IsFileExist(targetPath) {
		ctx.JSON(http.StatusNotFound, common.NotFound)
		return
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", imgName))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)
}

func (p *PostHandler) GetPostById(ctx *gin.Context) {

}


func (p *PostHandler) CreatePost(ctx *gin.Context) {
	var newPost models.Post

	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		ctx.JSON(http.StatusBadRequest, common.RespErr{Message: common.EmptyFieldErr})
		return
	}
	post, err := p.PostUSeCase.CreatePost(newPost)
	if err != nil {
		ctx.JSON(err.StatusCode(), err.Msg())
		return
	}

	ctx.JSON(http.StatusOK, post)
}
