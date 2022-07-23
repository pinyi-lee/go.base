package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}

func Rand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	if max-min > 0 {
		return rand.Intn(max-min) + min
	}

	return min
}

func Rand64(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())

	if max-min > 0 {
		return rand.Int63n(max-min) + min
	}

	return min
}

func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func Timestamp() int64 {
	return time.Now().UnixMilli()
}

func StructToJsonString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
