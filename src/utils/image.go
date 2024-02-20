package utils

import (
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

const length = 8
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateNewFilename() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetFileExtension(path string) string {
	return strings.ToLower(filepath.Ext(path))
}
