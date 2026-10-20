package parser

import (
	"bufio"
	"http_server1/engine/constants"
	"http_server1/engine/model"
	"io"
	"net"
	"slices"
	"strconv"
	"strings"
)

func ParseRequest(conn net.Conn) model.Request {
	return parseNewRequestFromConnection(conn)
}

func parseNewRequestFromConnection(conn net.Conn) model.Request {
	reader := bufio.NewReader(conn)
	lines := make(map[string]string)
	content := ""
	contentLength := 0

	method, target, version := parseRequestLine(reader)

	for {
		rawLine, _ := reader.ReadSlice(constants.LF)
		line := rawLine
		if line[0] == constants.CR {
			line = rawLine[1:]
		}

		if isEndOfStream(line) {
			break
		}

		key, value := parseLine(line)
		lines[key] = value
		if strings.Compare(key, constants.CONTENT_LENGTH_HEADER) == 0 {
			contentLength, _ = strconv.Atoi(value)
		}
	}

	if contentLength != 0 {
		content = parseContent(conn, contentLength)
	}

	return model.NewRequest(lines, method, target, version, content, contentLength)
}

func isEndOfStream(line []byte) bool {
	return len(line) > 0 &&
		(line[0] == constants.LF)
}

func parseRequestLine(reader *bufio.Reader) (string, string, string) {
	line, _, _ := reader.ReadLine()
	split := strings.Split(string(line), " ")
	return split[0], split[1], split[2]
}

func parseLine(line []byte) (string, string) {
	index := slices.Index(line, constants.DELIMITER)
	return string(line[:index]), strings.TrimLeft(string(line[index+1:]), " ")
}

func parseContent(conn net.Conn, contentLength int) string {
	contentBytes := make([]byte, contentLength)
	io.ReadFull(conn, contentBytes) // TODO: error handling
	return string(contentBytes)
}
