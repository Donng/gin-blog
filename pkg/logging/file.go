package logging

import (
	"fmt"
	"gin-blog/pkg/file"
	"gin-blog/pkg/setting"
	"os"
	"time"
)

var (
	LogSavePath = setting.App.LogSavePath
	LogSaveName = setting.App.LogSaveName
	LogFileExt  = setting.App.LogFileExt
	TimeFormat  = setting.App.TimeFormat
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.App.RuntimeRootPath, setting.App.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.App.LogSaveName,
		time.Now().Format(setting.App.TimeFormat),
		setting.App.LogFileExt,
	)
}

func openLogFile(filename, filePath string) (*os.File, error) {
	// 获得当前运行的目录
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	open, err := file.Open(src+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return open, nil
}
