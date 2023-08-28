package server

import (
	"fmt"
	"net"
	"time"
	"strings"
	"strconv"
	"net/http"
)

func CheckUsedPort(defaultPort int) string {
	check, err := net.Listen("tcp", ":" + strconv.Itoa(defaultPort))
	if err != nil {
		return fmt.Sprintf("Address already in use, port %d\n", defaultPort)
	}
	check.Close()
	return fmt.Sprintf("Address is available, port %d\n", defaultPort)
}

func GetClientIP(r *http.Request) string {
	ip := strings.Split(r.RemoteAddr, ":")
	return ip[0]
}

func DateFormat() string {
	dateFormat := fmt.Sprintf("%d/%d/%d", time.Now().Day(), time.Now().Month(), time.Now().Year())
	return dateFormat
}

func TimeFormat() string {
	timeFormat := fmt.Sprintf("%d:%d:%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	return timeFormat
}

func ProtoMethod(r *http.Request) string {
	protoMethod := fmt.Sprintf(`"%s %s %s"`, r.Method, r.URL.Path, r.Proto)
	return protoMethod
}