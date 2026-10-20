package listener

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"http_server1/engine/constants"
)

func ListenForHttpTcp(listenMode constants.ListenMode, sink chan<- byte) {
	var logPrefix string = fmt.Sprintf("listener.listenForHttpTcp mode \"%s\" :: ", listenMode.Desc())
	log.SetPrefix(logPrefix)

	var listenerPort string = ":" + strconv.Itoa(constants.HTTP_PORT)
	networkCardListener, err := net.Listen(constants.HTTP_TCP, listenerPort)
	if err != nil {
		log.Fatal("error creating a listener: " + err.Error())
	}
	defer networkCardListener.Close()

	for {
		connection, err := networkCardListener.Accept()
		if err != nil {
			log.Println("cannot accept connection: " + connection.LocalAddr().String())
		}
		log.Println("connected: " + connection.RemoteAddr().String())
		go func(conn net.Conn) {
			conn.SetReadDeadline(time.Now().Add(1000 * time.Millisecond))
			var incomingBytes []byte = make([]byte, 1024)
			n, err := conn.Read(incomingBytes)
			if err != nil {
				log.Println("error reading from connection")
			}
			fmt.Println("--------")
			fmt.Println(string(incomingBytes))
			fmt.Printf("Read %d bytes\n", n)
			fmt.Println("--------")

			conn.Write([]byte("sample server"))
			conn.Close()
		}(connection)
	}
}
