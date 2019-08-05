package config

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Unknwon/goconfig"
)

var (
	CFG  *goconfig.ConfigFile
	ROOT string
)

// fileExist 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Load(fileName string, moreFiles ...string) error {
	//合成参数为数组
	fileNames := make([]string, 1, len(moreFiles)+1)
	fileNames[0] = fileName
	if len(moreFiles) > 0 {
		fileNames = append(fileNames, moreFiles...)
	}

	//获取exe执行路径
	curFilename := os.Args[0]
	binaryPath, err := exec.LookPath(curFilename)
	if err != nil {
		return errors.New("binary path error")
	}

	binaryPath, err = filepath.Abs(binaryPath)
	if err != nil {
		return errors.New("binary abs path error")
	}
	ROOT = filepath.Dir(binaryPath)

	//拼接
	newConfigList := []string{}
	for _, name := range fileNames {
		new_name := ROOT + name
		//go run 时路径在运行环境，所以需要校正
		if !fileExist(new_name) {
			curDir, _ := os.Getwd()
			pos := strings.LastIndex(curDir, "src")
			if pos == -1 {
				return errors.New("can't find " + new_name)
			}

			ROOT = curDir[:pos]
			new_name = ROOT + name
		}
		newConfigList = append(newConfigList, new_name)
	}

	CFG, err = goconfig.LoadConfigFile(newConfigList[0], newConfigList[1:]...)
	if err != nil {
		return err
	}

	return nil

}
