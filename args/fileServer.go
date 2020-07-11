/*
 * @Author: your name
 * @Date: 2020-07-11 12:03:54
 * @LastEditTime: 2020-07-11 15:43:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/args/fileServer.go
 */

package args

import (
	"MyDiskServer/conf"
	"os"
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
func (fp *FilePath) CheckRequestPath(path string) (flag bool) {
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
	if len(targetPath) > len(rootPathList) {
		var headPath = strings.Join(targetPath[0:len(rootPathList)], "/")
		if headPath != strings.Join(rootPathList, "/") {
			flag = false
		}
	} else {
		flag = false
	}
	flag = true
	return
}

// GetNewPath 获取新的路径
func (fp *FilePath) GetNewPath(path string) (newPath string) {
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

// GetResult 获取请求结果
func (fp *FilePath) GetResult(fileInfoList []os.FileInfo) (result []map[string]interface{}) {
	result = make([]map[string]interface{}, 0)
	for _, v := range fileInfoList {
		var filename = v.Name()
		var fileType = v.IsDir()
		var filePerm = v.Mode().String()
		var permMap = map[string]string{
			"own":   "",
			"group": "",
			"other": "",
		}
		var mark = 0
		for i := range filePerm[1:] {
			var perm string
			switch filePerm[1:][i : i+1] {
			case "-":
				perm = "0"
			default:
				perm = "1"
			}
			switch {
			case mark <= 2:
				permMap["own"] += perm
			case mark >= 2 && mark <= 5:
				permMap["group"] += perm
			default:
				permMap["other"] += perm
			}
			mark++
		}
		result = append(result, map[string]interface{}{
			"filename": filename,
			"filetype": fileType,
			"filePerm": permMap,
		})
	}
	return
}

// RenameFile 重命名文件
type RenameFile struct {
	OldPath string
	NewFile string
}

// GetRequestPath 获取请求的目标路径
func (rf *RenameFile) GetRequestPath() (path string) {
	switch conf.Path {
	case "/":
		path = rf.OldPath
	default:
		path = conf.Path + rf.OldPath
	}
	return
}

// CheckPath 检查路径是否合规
func (rf *RenameFile) CheckPath(path string) (flag bool) {
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
	if len(targetPath) > len(rootPathList) {
		var headPath = strings.Join(targetPath[0:len(rootPathList)], "/")
		if headPath != strings.Join(rootPathList, "/") {
			flag = false
		}
	} else {
		flag = false
	}
	flag = true
	return
}
