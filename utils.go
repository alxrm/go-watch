package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func md5By(filePath string) (string, error) {
	file, errOpen := os.Open(filePath)

	var result string

	if errOpen != nil {
		return result, errOpen
	}

	defer file.Close()

	hash := md5.New()
	_, errCopy := io.Copy(hash, file)

	if errCopy != nil {
		return result, errCopy
	}

	result = hex.EncodeToString(hash.Sum(nil))

	return result, nil
}

func validateHash(hash string) bool {
	log.Println(hash)
	return !(strings.Contains(hash, "/") || strings.Contains(hash, ".")) && len(hash) > 0
}
