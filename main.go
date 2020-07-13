/*
 * @Author: xiaoboya
 * @Date: 2020-06-19 16:15:26
 * @LastEditTime: 2020-07-13 15:32:09
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/main.go
 */

package main

import "MyDiskServer/core"

func main() {
	go core.InitServer()
	go core.InitProxy()
	for {
	}
}
