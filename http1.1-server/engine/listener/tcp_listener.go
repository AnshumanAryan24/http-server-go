package listener

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"http_server1/engine/constants"
	"http_server1/engine/parser"
)

func ListenForHttpTcp(listenMode constants.ListenMode, sink chan<- byte) {
	var logPrefix string = fmt.Sprintf("listener.listenForHttpTcp mode \"%s\" :: ", listenMode.Desc())

	var listenerPort string = ":" + strconv.Itoa(constants.HTTP_PORT)
	networkCardListener, err := net.Listen(constants.HTTP_TCP, listenerPort)
	if err != nil {
		log.Fatal(logPrefix + "error creating a listener: " + err.Error())
	}
	defer networkCardListener.Close()

	for {
		connection, err := networkCardListener.Accept()
		if err != nil {
			log.Println(logPrefix + "cannot accept connection: " + connection.LocalAddr().String())
		}
		log.Println(logPrefix + "connected: " + connection.RemoteAddr().String())
		go func(conn net.Conn) {
			conn.SetDeadline(time.Now().Add(time.Millisecond * 200))
			request := parser.ParseRequest(conn)
			fmt.Println(request.String())
			writeByteCount, writeErr := conn.Write([]byte("HTTP/1.1 200 OK" +
				"\r\n" +
				constants.CONTENT_LENGTH_HEADER + string(constants.DELIMITER) + "11" +
				"\r\n" +
				"\r\n" +
				"Hello World"))
			conn.Close()
			fmt.Printf("Written %d bytes with error:\n%v\n\n", writeByteCount, writeErr)
		}(connection)
	}
}
