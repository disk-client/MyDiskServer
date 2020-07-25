/*
 * @Author: your name
 * @Date: 2020-07-03 22:07:35
 * @LastEditTime: 2020-07-25 17:20:00
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskServer/core/initConnect.go
 */

package core

import (
	"MyDiskServer/conf"
	"MyDiskServer/utils"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func connectControl(username string) {
	var tcpAddr *net.TCPAddr
	//这里在一台机测试，所以没有连接到公网，可以修改到公网ip
	tcpAddr, _ = net.ResolveTCPAddr("tcp", conf.HOST+":"+conf.HEARTPORT)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}
	var content = "SERVER" + time.Now().Format("2006-01-02 15:04:05") + "||" + username
	encrypted, err := utils.AesEncrypt([]byte(content), conf.AesKey)
	if err != nil {
		fmt.Println("encrypt error ! " + err.Error())
		return
	}
	conn.Write([]byte(encrypted))
	fmt.Println(conn.LocalAddr().String() + " : Client connected!8089")
	reader := bufio.NewReader(conn)
	for {
		s, err := reader.ReadString('\n')
		switch err {
		case io.EOF:
			continue
		case nil:
			switch s {
			case "new\n":
				go combine()
			case "hi\n":
				fmt.Println("保持链接中...")
			case "fuck\n":
				os.Exit(200)
			}
		default:
			break
		}
	}
}

func combine() {
	local := connectLocal()
	remote := connectRemote()
	if local != nil && remote != nil {
		joinConn(local, remote)
	} else {
		if local != nil {
			err := local.Close()
			if err != nil {
				fmt.Println("close local:" + err.Error())
			}
		}
		if remote != nil {
			err := remote.Close()
			if err != nil {
				fmt.Println("close remote:" + err.Error())
			}

		}
	}
}

// 对接两个端数据
func joinConn(local *net.TCPConn, remote *net.TCPConn) {
	f := func(local *net.TCPConn, remote *net.TCPConn) {
		defer local.Close()
		defer remote.Close()
		_, err := io.Copy(local, remote)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("end")
	}
	go f(local, remote)
	go f(remote, local)
}

// 代理本机的8000端口，也就是后台服务
func connectLocal() *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8000")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return nil
	}
	fmt.Println(conn.LocalAddr().String() + " : Client connected!8000")
	return conn
}

// 链接远端的8088端口
func connectRemote() *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", conf.HOST+":"+conf.TUNNELPORT)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return nil
	}
	fmt.Println(conn.LocalAddr().String() + " : Client connected!8088")
	return conn
}

// InitProxy 初始化代理
func InitProxy() {
	connectControl("xiaoboya")
}
