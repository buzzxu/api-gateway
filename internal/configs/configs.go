package configs

import (
	"fmt"
	"github.com/buzzxu/boys/common/files"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"time"
)

type YamlLoader struct {
	apps     *Apps
	services *Services
}
type MySQLLoader struct {
	app      *App
	services *Services
}

func (l YamlLoader) Load() {
	root := getCurrentAbPathByCaller()
	appFile := root + "/apps/app.yml"
	if !files.Exists(appFile) {
		panic(fmt.Sprintf("%s 不存在,请检查文件路径", appFile))
	}
	app, err := ioutil.ReadFile(appFile)
	if err != nil {
		panic(fmt.Sprintf("读取文件: %s 发生错误,原因: %s", appFile, err.Error()))
	}
	err = yaml.Unmarshal(app, &l.apps)
	if err != nil {
		panic(fmt.Sprintf("文件解析失败: %s", err.Error()))
		os.Exit(1)
	}
	serviceFile := root + "/services.yml"
	if !files.Exists(serviceFile) {
		panic(fmt.Sprintf("%s 不存在,请检查文件路径", appFile))
	}
	services, err := ioutil.ReadFile(serviceFile)
	if err != nil {
		panic(fmt.Sprintf("读取文件: %s 发生错误,原因: %s", appFile, err.Error()))
	}
	err = yaml.Unmarshal(services, &l.services)
	if err != nil {
		panic(fmt.Sprintf("文件解析失败: %s", err.Error()))
		os.Exit(1)
	}
}

func (l YamlLoader) Apps() *Apps {
	return l.apps
}

func (l YamlLoader) Services() *Services {
	return l.services
}

func (l YamlLoader) Refresh(duration time.Duration) {

}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
