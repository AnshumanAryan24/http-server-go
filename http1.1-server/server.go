package main

import (
	"fmt"
	"http_server1/engine/constants"
	"http_server1/engine/listener"
)

func main() {
	listenMode, _ := constants.NewListenMode(constants.ACCEPT_ALL)
	fmt.Println("HTTP 1.1 server")
	listener.ListenForHttpTcp(listenMode, nil)
}
