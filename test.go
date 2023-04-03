package main

import (
	"github.com/helogic/frp/cmd/frpc"
	"os"
)

func main() {
	//os.Getenv("server_addr")
	//os.Getenv("server_port")
	//os.Getenv("local_port")
	//os.Getenv("local_ip")
	os.Setenv("frp_server_addr", "127.0.0.1")
	os.Setenv("frp_server_port", "7000")
	os.Setenv("frp_local_ip", "127.0.0.1")
	os.Setenv("frp_local_port", "80")
	frpc.Run()
}
