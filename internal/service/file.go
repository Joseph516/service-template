package service

import (
	"errors"
	"mime/multipart"
	"os"
	"path"
	"service-template/global"
	"service-template/internal/model"
	"service-template/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

type GetFileRequest struct {
	Filename    string `json:"filename"`
	Type        int    `json:"type"`
	CreatedTime int64  `json:"created_time"`
	CreatedBy   string `json:"created_by"`
}

/*
UploadFile 文件上传
Inputs:
	- fileType: 文件类型
	- file: 表单文件数据
	- fileHeader： 文件头部
	- username：用户名
Outputs:
	- FileInfo: 文件信息
	- error: 错误内容，若成功则返回nil
*/
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader, username string) (*FileInfo, error) {
	// fileName := upload.GetFileNameMD5(fileHeader.Filename)
	fileName := fileHeader.Filename // 不对文件名加密
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := path.Join(upload.GetSavePath(), username)
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	// 保存文件
	// 重名检查，如果存在重名，则直接在文件名后加数字"(n)"
	fileName = svc.dao.CheckDuplicatedFilename(fileName, username)

	dst := path.Join(uploadSavePath, fileName)
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := path.Join(global.AppSetting.UploadServerUrl, username, fileName)

	// 将文件上传记录写入数据库
	err := svc.dao.UploadFile(fileName, "", username, accessUrl, int(fileType))
	if err != nil {
		return nil, err
	}

	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

// GetUploadFile 根据请求参数返回数据中文件信息
func (svc *Service) GetUploadFile(params *GetFileRequest) (*[]model.File, error) {
	// TODO
	svc.dao.GetUploadFile()
	return nil, nil
}
