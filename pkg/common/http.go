package common

import (
	"net/http"
	"strings"
)

func GetCurrentIP(r http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
