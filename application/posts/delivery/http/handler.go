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
	PostUseCase    posts.UseCase
	SessionBuilder common.SessionBuilder
}

type RespImg struct {
	Location string `json:"location"`
}

func NewRest(router *gin.RouterGroup, useCase posts.UseCase, sessionBuilder common.SessionBuilder, AuthRequired gin.HandlerFunc, isAdminOrMethodist gin.HandlerFunc) *PostHandler {
	rest := &PostHandler{PostUseCase: useCase, SessionBuilder: sessionBuilder}
	rest.routes(router, AuthRequired, isAdminOrMethodist)
	return rest
}

func (p *PostHandler) routes(router *gin.RouterGroup, AuthRequired gin.HandlerFunc, isAdminOrMethodist gin.HandlerFunc) {
	router.GET("/current-id", p.GetCurrentPostId)
	router.POST("/:post-id/upload-image", p.UploadPostsImages)
	router.POST("/", p.CreatePost)
	router.GET("/:post-id/download-image/:image-name", p.DownloadPostImage)
	router.GET("/:post-id", p.GetPostById)
	router.GET("/list", p.GetPosts)
	router.DELETE("/:post-id", p.DeletePostById)
	router.Use(AuthRequired)
	{
	}
}

// GetCurrentPostId
// @Summary GetCurrentPostId
// @Description Returns a free id for post in the database.
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success 200 {integer} integer
// @Failure 500 {object} common.RespErr
// @Router /posts/current-id [get]
func (p *PostHandler) GetCurrentPostId(ctx *gin.Context) {
	id, err := p.PostUseCase.GetCurrentPostId()
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

// GetPostById
// @Summary GetPostById
// @Description Returns a post by proceed id.
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param post-id path int true "Post Id"
// @Success 200 {object} models.Post
// @Failure 400 {object} common.BadReqErr
// @Failure 404 {object} common.NotFoundErr
// @Failure 500 {object} common.RespErr
// @Router /posts/{post-id} [get]
func (p *PostHandler) GetPostById(ctx *gin.Context) {
	var req struct {
		PostId int `uri:"post-id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.EmptyFieldErr)
		return
	}
	post, err := p.PostUseCase.GetPostById(req.PostId)
	if err != nil {
		ctx.JSON(err.StatusCode(), err.Msg())
		return
	}
	ctx.JSON(http.StatusOK, post)
}

type reqPostCreate = struct {
	AuthorId int             `json:"author_id" example:"1"`
	TagType  string          `json:"tag_type" example:"general|important|education"`
	Title    string          `json:"title" example:"Bla-Bla Post"`
	Content  string          `json:"content" example:"bla-bla-bla"`
}

// CreatePost
// @Summary CreatePost
// @Description Create post with proceed data.
// @Accept  json
// @Produce  json
// @Param 	data body reqPostCreate true "Post data"
// @Success 200 {object} models.Post
// @Failure 400 {object} common.BadReqErr
// @Failure 500 {object} common.RespErr
// @Router /posts [post]
func (p *PostHandler) CreatePost(ctx *gin.Context) {
	var newPost models.Post

	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		ctx.JSON(http.StatusBadRequest, common.RespErr{Message: common.EmptyFieldErr})
		return
	}
	post, err := p.PostUseCase.CreatePost(newPost)
	if err != nil {
		ctx.JSON(err.StatusCode(), err.Msg())
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// GetPosts
// @Summary GetPosts
// @Description Returns certain number of posts entries if there are start and limit params. Otherwise returns all entries.
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param  start query int false "start of output of records"
// @Param  limit query int false "limit of output of records"
// @Success 200 {array} models.Post
// @Failure 400 {object} common.BadReqErr
// @Failure 500 {object} common.RespErr
// @Router /posts/list [get]
func (p *PostHandler) GetPosts(ctx *gin.Context) {
	var req struct {
		Start uint16 `form:"start"`
		Limit uint16 `form:"limit"`
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.EmptyFieldErr)
		return
	}

	postsList, err := p.PostUseCase.GetPostsList(req.Start, req.Limit)
	if err != nil {
		ctx.JSON(err.StatusCode(), err.Msg())
		return
	}
	ctx.JSON(http.StatusOK, postsList)
}

// DeletePostById
// @Summary DeletePost
// @Description Delete post by proceed id.
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param post-id path int true "Post Id"
// @Success 200
// @Failure 400 {object} common.BadReqErr
// @Failure 404 {object} common.NotFoundErr
// @Failure 500 {object} common.RespErr
// @Router /posts/{post-id} [delete]
func (p *PostHandler) DeletePostById(ctx *gin.Context) {
	var req struct {
		PostId int `uri:"post-id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.EmptyFieldErr)
		return
	}
	err := p.PostUseCase.DeletePost(req.PostId)
	if err != nil {
		ctx.JSON(err.StatusCode(), err.Msg())
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
