package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("test")
	startip := net.ParseIP("192.168.100.0")
	endip := net.ParseIP("192.168.100.14")
	println(compIp(startip, endip))
}

func compIp(startip net.IP, endip net.IP) int {
	if ok := startip.Equal(endip); ok {
		return 0
	}
	for i, ip := range startip {
		if ip > endip[i] {
			return 1
		}
	}
	return -1
}
