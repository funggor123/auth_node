package common

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func GetCurrentTimeInRFC() string {
	return time.Now().Format("2006-01-02T15:04:05")
}

func GetLinuxTime() int32 {
	return int32(time.Now().Unix())
}

func GetTimeInAllInt(t time.Time) string {
	return t.Format("20060102150405")
}

func GetUTCTimeInAllInt(day int) string {
	dd, _ := time.ParseDuration(strconv.Itoa(day*24) + "h")
	return time.Now().Add(dd).Format("20060102150405")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetCurrentProgramPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return dir, err
	}
	return dir, nil
}
