/*
package api upload 管理文件的上传功能
*/

package api

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"service-template/internal/service"
	"service-template/pkg/app"
	"service-template/pkg/errcode"
	"service-template/pkg/upload"
	"strconv"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) CheckFile(c *gin.Context) (multipart.File, *multipart.FileHeader, int) {
	file, fileHeader, err := c.Request.FormFile("file")
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return nil, nil, 0
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		response.ToErrResponse(errcode.InvalidParams.WithDetails("Invalid file type."))
		return nil, nil, 0
	}
	return file, fileHeader, fileType
}

// UploadFile 文件上传
func (u Upload) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		response.ToErrResponse(errcode.InvalidParams.WithDetails("Invalid file type."))
		return
	}
	// 从token中获取当前用户
	token := c.GetHeader("token")
	claims, err := app.ParseToken(token)
	if err != nil {
		response.ToErrResponse(errcode.UnauthorizedTokenError.WithDetails(err.Error()))
		return
	}
	// 将文件保存到指定位置
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader, claims.AppKey)
	if err != nil {
		response.ToErrResponse(errcode.UploadFileFailed.WithDetails(err.Error()))
		return
	}
	// 上传成功
	data := gin.H{
		"file_access_url": fileInfo.AccessUrl,
	}
	response.ToResponse(data)
}

// GetFile 获取文件
func (u Upload) GetFile(c *gin.Context) {
	param := service.GetFileRequest{}
	response := app.NewResponse(c)

	// 参数绑定与校验
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	files, err := svc.GetUploadFile(&param)

	if err != nil {
		response.ToErrResponse(errcode.RegisterFailed.WithDetails(err.Error()))
		return
	}
	response.ToResponse(files)
}

func (u Upload) List(c *gin.Context) {
	// TBD
}
