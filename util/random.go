package util

import (
	"math/rand"
	"strings"
	"time"
)

var random *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + random.Int63n(max-min+1)
}

func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[random.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomFileType() string {
	return RandomString(6)
}

func RandomExt() string {
	return RandomString(3)
}

func RandomFolderName() string {
	return RandomString(8)
}
