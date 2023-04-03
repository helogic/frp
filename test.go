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
	os.Setenv("server_addr", "127.0.0.1")
	os.Setenv("server_port", "7000")
	os.Setenv("local_ip", "127.0.0.1")
	os.Setenv("local_port", "80")
	frpc.Run()
}
