/*
 * @Author: your name
 * @Date: 2020-06-19 18:27:04
 * @LastEditTime: 2020-06-20 09:39:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/utils/aboutUDP.go
 */

package utils

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"gopkg.in/ini.v1"
)

// GetPlatformAddr 和云端平台保持心跳
func GetPlatformAddr() (host string, port int) {
	iniObj, err := ini.Load("../conf/admin.ini")
	PanicErr(err)
	var sec = iniObj.Section("platform")
	host = sec.Key("HOST").String()
	var portStr = sec.Key("PORT").String()
	port, err = strconv.Atoi(portStr)
	PanicErr(err)
	return
}

// HeartBeat 平台心跳
func HeartBeat() {
	var host, port = GetPlatformAddr()
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 9202} // 注意端口必须固定
	dstAddr := &net.UDPAddr{IP: net.ParseIP(host), Port: port}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	// if _, err = conn.Write([]byte("s" + conf.TheUser.Name)); err != nil {
	if _, err = conn.Write([]byte("s" + "xiaoboya")); err != nil {
		log.Panic(err)
	}
	conn.Close()
}

// KeepHeart 保持平台心跳
func KeepHeart() {
	for {
		HeartBeat()
		time.Sleep(10 * time.Second)
	}
}
