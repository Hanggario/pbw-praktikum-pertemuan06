package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("GarinHanggario", 0755) // Membuat direktori dengan izin 0755 (rwxr-xr-x)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'GarinHanggario' berhasil dibuat.")
}
