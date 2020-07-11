/*
 * @Author: 肖博雅
 * @Date: 2020-07-11 13:06:34
 * @LastEditTime: 2020-07-11 21:23:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/core/fileServer.GO
 */

package core

import (
	"MyDiskServer/args"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// InitServer 初始化服务
func InitServer() {
	var r = gin.Default()
	r.POST("/FileMenu", getDirList)
	r.POST("/NewDir", newDir)
	r.POST("/Rename", rename)
	r.POST("/Download", download)
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
	var newPath = info.GetNewPath(path)
	fileInfoList, err := ioutil.ReadDir(newPath)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "读取文件失败",
		})
		return
	}
	var result = info.GetResult(fileInfoList)
	c.JSON(200, gin.H{
		"msg":  "返回成功",
		"data": result,
	})
}

func newDir(c *gin.Context) {
	var info args.FilePath
	var err = c.BindJSON(&info)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"msg":  "参数异常",
			"succ": false,
		})
		return
	}
	if ok := info.CheckRequestPath(info.Path); !ok {
		c.JSON(200, gin.H{
			"msg":  "当前路径错误，无新建文件夹权限",
			"succ": false,
		})
		return
	}
	err = os.Mkdir(info.Path, 1660)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":  "新建文件夹失败",
			"succ": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "新建文件夹成功",
		"succ": true,
	})
	return
}

func rename(c *gin.Context) {
	var info args.RenameFile
	var err = c.BindJSON(&info)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"msg": "参数异常",
		})
		return
	}
	var path = info.GetRequestPath()
	if ok := info.CheckPath(path); !ok {
		c.JSON(200, gin.H{
			"msg": "当前路径错误，无新建文件夹权限",
		})
		return
	}
	var pathList = strings.Split(path, "/")
	pathList[len(pathList)-1] = info.NewFile
	var newName = strings.Join(pathList, "/")
	err = os.Rename(path, newName)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "重命名失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "重命名成功",
	})
	return
}

func download(c *gin.Context) {
	var info args.FilePath
	var err = c.BindJSON(&info)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "参数异常",
		})
		return
	}
	var path = info.GetRequestPath()
	if ok := info.CheckRequestPath(path); !ok {
		c.JSON(200, gin.H{
			"msg": "当前路径错误，无新建文件夹权限",
		})
	} else {
		var newPath = info.GetNewPath(path)
		c.File(newPath)
	}
	return
}
