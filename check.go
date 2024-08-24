package main

import (
	"fmt"
	"net"
	"time"
)

func Check(host string, port string) string {
	address := host + ":" + port
	timeout := time.Duration(time.Second * 10)
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN]-%v is not responding. \n Error: %v\n", address, err)
	} else {
		status = fmt.Sprintf("[UP]-%v is responding. \n From: %v To: %v\n", address, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
