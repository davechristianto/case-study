package main

import (
	"fmt"
	"os"
)

func main() {
	var hours, minutes, seconds int

	// Contoh Input 12:00:00
	fmt.Print("Masukkan waktu di Bumi (format HH:MM:SS): ")
	_, err := fmt.Scanf("%d:%d:%d", &hours, &minutes, &seconds)
	if err != nil {
		fmt.Println("Format input salah. Contoh: 12:00:00")
		os.Exit(1)
	}

	totalEarthSeconds := hours*3600 + minutes*60 + seconds

	totalRoketinSeconds := totalEarthSeconds * 100000 / 86400

	roketinHours := totalRoketinSeconds / (100 * 100)
	sisa := totalRoketinSeconds % (100 * 100)
	roketinMinutes := sisa / 100
	roketinSeconds := sisa % 100

	fmt.Printf("Di Bumi %02d:%02d:%02d, di Planet Roketin %02d:%02d:%02d\n",
		hours, minutes, seconds, roketinHours, roketinMinutes, roketinSeconds)
}