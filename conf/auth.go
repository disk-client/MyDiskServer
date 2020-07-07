/*
 * @Author: your name
 * @Date: 2020-06-19 18:42:12
 * @LastEditTime: 2020-07-07 17:27:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/conf/auth.go
 */

package conf

// User 用户
type User struct {
	Name string
}

// TheUser 当前用户
var TheUser User

// Addr 地址
var Addr = ""
