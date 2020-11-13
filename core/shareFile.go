/*
 * @Author: your name
 * @Date: 2020-07-29 21:10:32
 * @LastEditTime: 2020-07-30 10:11:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/core/shareFile.go
 */

package core

import (
	"MyDiskServer/conf"
	"MyDiskServer/utils"
	"encoding/json"
	"fmt"
	"os"
)

// FileDownloadCert 文件下载证书
type FileDownloadCert struct {
	Path string `json:"path"`
	Days int    `json:"days"`
	User string `json:"user"`
}

func (cert *FileDownloadCert) check() (ok bool) {
	ok = false
	if cert.Days < 0 {
		return
	}
	if cert.Path == "" {
		return
	}
	if pathOK := utils.CheckPath(cert.Path); !pathOK {
		fmt.Println(1111)
		return
	}
	fileInfo, err := os.Stat(cert.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
	}
	if isDir := fileInfo.IsDir(); isDir {
		return
	}
	return true
}

// ProduceCert 产生文件
func ProduceCert(days int, path string) (result string) {
	result = ""
	var cert = FileDownloadCert{Path: path, Days: days, User: conf.TheUser.Name}
	if pathOK := cert.check(); !pathOK {
		return "路径错误"
	}
	resultBytes, err := json.Marshal(cert)
	if err != nil {
		return "序列化错误"
	}
	resultBytes, err = utils.AesEncrypt(resultBytes, conf.CryptCertKey)
	if err != nil {
		return "加解密错误"
	}
	result = string(resultBytes)
	return
}
