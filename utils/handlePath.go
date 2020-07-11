/*
 * @Author: your name
 * @Date: 2020-06-19 17:54:52
 * @LastEditTime: 2020-07-11 13:08:39
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/utils/handlePath.go
 */

package utils

import (
	"strings"
)

// RootPath 根路径
var RootPath = ""

// NowPath 当前路径
var NowPath = ""

// CheckPath 检查路径是否合规
func CheckPath(path string) bool {
	var NowPathBak = NowPath
	var pathList = strings.Split(path, "/")
	for _, info := range pathList {
		switch info {
		case "", ".":
			NowPathBak = NowPathBak + ""
		case "..":
			var p = strings.Split(NowPathBak, "/")
			if len(p) != 0 {
				p = p[0 : len(p)-1]
				NowPathBak = strings.Join(p, "/")
			}
			return false
		default:
			NowPathBak += "/" + info
		}
	}
	if len(NowPathBak) < len(NowPath) {
		return false
	}
	if NowPathBak[0:len(NowPath)] != NowPath {
		return false
	}
	NowPath = NowPathBak
	return true
}
