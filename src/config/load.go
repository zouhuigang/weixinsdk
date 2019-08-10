package config

import (
	"errors"
	"fmt"
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

// fileExist 检查文件或目录是否存在,如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//核对文件路径
func checkFileExist(name string) (new_name string, err error) {
	//go run 时路径在运行环境，所以需要校正
	new_name = ROOT + name
	if !fileExist(new_name) {
		curDir, _ := os.Getwd()
		pos := strings.LastIndex(curDir, "src")
		if pos == -1 {
			return new_name, errors.New("can't find " + new_name)
		}

		ROOT = curDir[:pos]
		new_name = ROOT + name
	}

	return new_name, nil
}

//加载多个其他文件
func getMoreFiles(moreFiles ...string) ([]string, error) {
	//合成参数为数组
	fileNames := make([]string, 0, len(moreFiles))
	if len(moreFiles) > 0 {
		fileNames = append(fileNames, moreFiles...)
	}
	//拼接
	newConfigList := []string{}
	for _, name := range fileNames {
		new_name, err := checkFileExist(name)
		if err != nil {
			return newConfigList, err
		}
		newConfigList = append(newConfigList, new_name)
	}

	return newConfigList, nil
}

//"/build/base.env.ini", "/build/dev.env.ini"
func Load() error {

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

	//加载base文件
	baseFile := `/build/base.env.ini`
	baseFile, err = checkFileExist(baseFile)
	if err != nil {
		return err
	}
	CFG, err = goconfig.LoadConfigFile(baseFile)
	if err != nil {
		return err
	}

	//额外文件
	env := CFG.MustValue("parameter", "env", "")
	envFile := fmt.Sprintf("/build/%s.env.ini", env)
	moreFiles, err := getMoreFiles(envFile)
	if err != nil {
		return err
	}
	err = CFG.AppendFiles(moreFiles...)
	if err != nil {
		return err
	}

	return nil

}
