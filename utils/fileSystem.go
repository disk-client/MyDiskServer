/*
 * @Author: your name
 * @Date: 2020-06-19 17:37:03
 * @LastEditTime: 2020-06-19 17:54:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/utils/fileSystem.go
 */

package utils

import "io/ioutil"

// GetFileList 获取文件夹结构
func GetFileList(path string) interface{} {
	var fileInfoList, err = ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	var result = make([]map[string]interface{}, 0)
	for _, fileInfo := range fileInfoList {
		var filename = fileInfo.Name()
		var fileType = fileInfo.IsDir()
		var filePerm = fileInfo.Mode().String()
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
	return result
}
