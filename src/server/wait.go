package server

import (
	"fmt"
	"net"
	"time"
)

func WaitReady(ipAddress string, port int) {
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ipAddress, port), time.Second)
		if err, ok := err.(*net.OpError); ok && err.Timeout() {
			continue
		}
		if err != nil {
			continue
		}
		conn.Close()
		return
	}
}
