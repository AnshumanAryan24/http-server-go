package model

import (
	"errors"
	"http_server1/engine/constants"
	"strings"
)

type Request interface {
	Get(string) (string, error)
	GetContent() string
	GetContentLength() int
	String() string
}

type request struct {
	lines                            map[string]string
	method, target, version, content string
	contentLength                    int
}

func NewRequest(lines map[string]string, method, target, version, content string, contentLength int) Request {
	return &request{lines, method, target, version, content, contentLength}
}

func (r *request) Get(header string) (string, error) {
	var headerLine string = r.lines[header]
	if headerLine == "" {
		return "", errors.New("no such header: " + header)
	}
	return string(headerLine), nil
}

func (r *request) GetContentLength() int {
	return r.contentLength
}

func (r *request) GetContent() string {
	return r.content
}

func (r *request) String() string {
	var builder strings.Builder
	builder.WriteString(r.method)
	builder.WriteString(" ")
	builder.WriteString(r.target)
	builder.WriteString(" ")
	builder.WriteString(r.version)
	builder.WriteString("\r\n")
	for k, v := range r.lines {
		builder.WriteString(k)
		builder.WriteByte(constants.DELIMITER)
		builder.WriteString(v)
		builder.WriteString("\n")
	}
	builder.WriteString("\n")
	builder.WriteString(r.content)
	return builder.String()
}
