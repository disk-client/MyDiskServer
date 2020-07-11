/*
 * @Author: your name
 * @Date: 2020-07-11 13:06:34
 * @LastEditTime: 2020-07-11 14:03:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/core/fileServer.GO
 */

package core

import (
	"MyDiskServer/args"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// InitServer 初始化服务
func InitServer() {
	var r = gin.Default()
	r.Run(":8000")
}

func getDirList(c *gin.Context) {
	var info args.FilePath
	var err = c.BindJSON(&info)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"msg": "参数异常",
		})
		return
	}
	var path = info.GetRequestPath()
	var newPath = info.CheckRequestPath(path)
	fileInfoList, err := ioutil.ReadDir(newPath)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "读取文件失败",
		})
		return
	}
	var result = make([]map[string]interface{}, 0)
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
	c.JSON(200, gin.H{
		"msg":  "返回成功",
		"data": result,
	})
}
