package common

import (
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

const (
	ImgDir           = "static"
	MaxImgSize       = 32 << 16 // 2 Mb
	MaxImgHeight     = 2500     //px
	MaxImgWidth      = 2500     //px
	UploadFileError  = -1
	//FileValid       = 0
	PngMime         = "image/png"
	JpegMime        = "image/jpeg"
	base64pngTitle  = 22
	base64jpegTitle = 23
	dirCreateErr    = "Не удалось создать директорию"
	someWentWrong   = "Что-то пошло не так, попробуйте позже."
)

func IsValidFile(header *multipart.FileHeader, allowedTypes []string) bool {
	mimeType := header.Header.Get("Content-type")
	for _, t := range allowedTypes {
		if t == mimeType {
			return true
		}
	}
	return false
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(dirname string) (string, Err) {
	//fileDir, _ := os.Getwd()
	pathDir := path.Join(PathToSaveStatic, dirname)

	if !IsFileExist(pathDir) {
		err := os.Mkdir(pathDir, 0777)
		if err != nil {
			return "", NewErr(http.StatusInternalServerError, dirCreateErr)
		}
	}

	return pathDir, nil
}

func AddOrUpdateUserFile(data io.Reader, imgPath string) (string, Err) {
	if data == nil {
		return "", nil
	}

	dst, err := os.Create(imgPath)
	if err != nil {
		return "", NewErr(http.StatusInternalServerError, someWentWrong)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, data); err != nil {
		return "", NewErr(UploadFileError, someWentWrong)
	}
	return imgPath, nil
}
