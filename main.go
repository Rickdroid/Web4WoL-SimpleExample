package main

import (
	"fmt"
	"net"
	"net/http"
)

//整理IP地址、MAC地址、魔术包内容等信息
type Message struct {
	Addr string
	Mac  []byte
}

func (m Message) getMessage() (string, []byte) {
	msg1 := Message{"ip:port", []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}}
	addr := msg1.Addr
	mac := msg1.Mac
	return addr, mac
}

func (m Message) getPackage() []byte {
	_, mac := m.getMessage()
	header := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	mpackage := append([]byte{}, header...)
	for i := 0; i < 16; i++ {
		mpackage = append(mpackage, mac...)
	}
	return mpackage
}

//魔术包的发送
func openPower(res http.ResponseWriter, req *http.Request) {
	//连接主机
	var m Message
	addr, _ := m.getMessage()
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//发送数据包
	conn.Write(m.getPackage())
}

//前台浏览器的监听和方法执行
func main() {
	http.HandleFunc("/openpower", openPower)
	http.ListenAndServe(":port", nil)
}
