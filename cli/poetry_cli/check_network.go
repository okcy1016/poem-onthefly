package main

import (
	"net"
	"time"
)

func ChkNetConn() {

	host := "mirrors.tuna.tsinghua.edu.cn"
	port := "443"
	timeout := time.Duration(5) * time.Second

	_, err := net.DialTimeout("tcp", (host + ":" + port), timeout)

	ChkErr(err)
}
