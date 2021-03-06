package upload

import (
	"fmt"
	"gin-demo/pkg/file"
	"gin-demo/pkg/setting"
	"gin-demo/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"gin-demo/pkg/logging"
)

func GetImageFullUrl (name string) string {
	return setting.AppConfig.ImagePrefixPath + "/" + GetImagePath() + name;
}

func GetImageName (name string) string{
	ext := path.Ext(name)
	filename := strings.TrimSuffix(name, ext)
	filename = util.EncodeMD5(filename)
	return filename + ext
}

func GetImagePath () string{
	return setting.AppConfig.ImageSavePath
}

func GetImageFullPath () string {
	return setting.AppConfig.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(filename string) bool {
	ext := file.GetExt(filename)
	for _, allowExt := range setting.AppConfig.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize (f multipart.File) bool {
	size , err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <= setting.AppConfig.ImagaMaxSize
}

func CheckImage (src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.GetWd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err !=  nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src err: %s", src)
	}
	return nil
}
