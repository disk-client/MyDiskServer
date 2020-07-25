/*
 * @Author: your name
 * @Date: 2020-07-07 17:43:50
 * @LastEditTime: 2020-07-25 17:19:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/conf/conf.go
 */

package conf

// HOST 远程主机
const HOST = "47.92.149.230"

// HEARTPORT 心跳端口
const HEARTPORT = "8089"

// TUNNELPORT 通道端口
const TUNNELPORT = "8088"

// Path 制定开始文件
var Path = ""

// AesKey 对称加密密钥
var AesKey = []byte("0f90023fc9ae101e")
