package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("GarinHanggario")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if fileInfo.IsDir() {
		fmt.Println("GarinHanggario adalah sebuah direktori.")
	} else {
		fmt.Println("GarinHanggario adalah sebuah file.")
	}
}
