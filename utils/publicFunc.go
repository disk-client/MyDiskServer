/*
 * @Author: your name
 * @Date: 2020-06-19 18:38:03
 * @LastEditTime: 2020-06-19 18:39:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/utils/publicFunc.go
 */

package utils

// PanicErr 异常打断程序
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
