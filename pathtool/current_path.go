package pathtool

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	RuntimePath string
)

func init() {
	RuntimePath = CurrentAbPath()
}

// CurrentAbPath 获取当前运行目录
func CurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string

	for i := 0; i <= 5; i++ {
		_, filename, _, ok := runtime.Caller(i)
		if ok && !strings.Contains(filename, "current_path") {
			abPath = filename
			break
		}
	}

	//_, filename, _, ok := runtime.Caller(2)
	//if ok {
	//	abPath = path.Dir(filename)
	//}

	abPath = path.Dir(abPath)

	return abPath
}
