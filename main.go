package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func randName(pfx, sfx string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return pfx + hex.EncodeToString(randBytes) + sfx
}

func create() {
	file := randName("filename-", ".txt") // file name
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}
	data := strings.Repeat("I hope you have a good day\n", 100) // message to fill files with
	if _, err := f.Write([]byte(data)); err != nil {
		f.Close()
		fmt.Print("Error: ", err.Error())
		return
	}
	if err := f.Close(); err != nil {
		fmt.Print("Error: ", err.Error())
		return
	}
}

func main() {
	num := 0
	amt := 3 // number of files to spam

	for num < amt {
		create()
		num++
	}
}
