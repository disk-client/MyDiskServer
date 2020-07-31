/*
 * @Author: your name
 * @Date: 2020-06-19 18:38:03
 * @LastEditTime: 2020-07-29 21:40:19
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/utils/publicFunc.go
 */

package utils

import (
	"fmt"
	"os"
)

// PanicErr 异常打断程序
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckFileIsExist 检查文件是否存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err)
			exist = false
		}
	}
	return exist
}
