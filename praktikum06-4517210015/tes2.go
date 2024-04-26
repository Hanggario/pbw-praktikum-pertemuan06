package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chmod("GarinHanggario", 0015)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Izin 'GarinHanggario' telah diubah menjadi 0015.")
}
