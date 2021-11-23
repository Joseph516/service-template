package dao

import (
	"fmt"
	"service-template/internal/model"
	"service-template/pkg/upload"
	"strconv"
	"time"
)

func (d *Dao) UploadFile(filename, desc, createdBy, accessUrl string, filetype int) error {
	file := model.File{
		Filename:      filename,
		Desc:          desc,
		CreatedBy:     createdBy,
		CreatedTime:   time.Now().Unix(),
		FileAccessUrl: accessUrl,
		Type:          filetype,
	}
	return file.Create(d.engine)
}

// CheckDuplicatedFilename 检查用户是否已经上传同名文件，若已经上传则直接对文件名增加数字编号
func (d Dao) CheckDuplicatedFilename(filename, username string) string {
	filenameOrigin, ext := upload.SparseFileName(filename)
	for i := 1; ; i += 1 {
		file := model.File{
			Filename:  filename,
			CreatedBy: username,
		}
		files, _ := file.GetFileByNameAuthor(d.engine)
		if len(files) > 0 {
			filename = fmt.Sprintf("%v(%v)%v", filenameOrigin, strconv.Itoa(i), ext)
		} else {
			return filename
		}
	}
}

func (d Dao) GetUploadFile()  {

}
