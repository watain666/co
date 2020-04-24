package os

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 当前操作系统是否Windows.
func IsWin() bool {
	return "windows" == runtime.GOOS
}

// 当前操作系统是否Linux.
func IsLinux() bool {
	return "linux" == runtime.GOOS
}

// 当前操作系统是否Mac OS/X.
func IsMac() bool {
	return "darwin" == runtime.GOOS
}

// 获取程序运行路径
func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

// 获取当前程序运行所在的路径,注意和Getwd有所不同.
func Pwd() string {
	var (
		dir, ex string
		err error
	)
	ex, err = os.Executable()
	if err == nil {
		exReal, _ := filepath.EvalSymlinks(ex)
		exReal, _ = filepath.Abs(exReal)
		dir = filepath.Dir(exReal)
	}
	return dir
}
