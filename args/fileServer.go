/*
 * @Author: your name
 * @Date: 2020-07-11 12:03:54
 * @LastEditTime: 2020-07-11 14:02:32
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/args/fileServer.go
 */

package args

import (
	"MyDiskServer/conf"
	"strings"
)

// FilePath 文件路径
type FilePath struct {
	Path string
}

// GetRequestPath 获取请求的目标路径
func (fp *FilePath) GetRequestPath() (path string) {
	switch conf.Path {
	case "/":
		path = fp.Path
	default:
		path = conf.Path + fp.Path
	}
	return
}

// CheckRequestPath 检查路径是否合理
func (fp *FilePath) CheckRequestPath(path string) (newPath string) {
	var rootPathList []string
	var targetPath = []string{""}
	var pathList = strings.Split(path, "/")
	for _, v := range pathList {
		if v == ".." {
			targetPath = targetPath[0 : len(targetPath)-1]
		} else if v == "." {
			continue
		} else {
			targetPath = append(targetPath, v)
		}
	}
	rootPathList = strings.Split(conf.Path, "/")
	newPath = strings.Join(targetPath, "/")
	if len(targetPath) > len(rootPathList) {
		var headPath = strings.Join(targetPath[0:len(rootPathList)], "/")
		if headPath != strings.Join(rootPathList, "/") {
			newPath = conf.Path
		}
	} else {
		newPath = conf.Path
	}
	return
}
