/*
 * @Author: xiaoboya
 * @Date: 2020-06-19 16:15:26
 * @LastEditTime: 2020-07-30 10:12:25
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/main.go
 */

package main

import (
	"MyDiskServer/core"
	"fmt"
)

func main() {
	// go core.InitServer()
	// go core.InitProxy()
	var cert = core.ProduceCert(3, "/Users/raymond/go/src/MyDiskServer/bin/mydisk_server_mac")
	fmt.Println(cert)
	for {
	}
}
