package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const server = "pool.ntp.org"

	currentTime, err := ntp.Time(server)
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Printf("Ошибка: %v", err)
		os.Exit(1)
	}

	fmt.Println("Точное текущее время:", currentTime.Local().Format(time.RFC3339Nano))
}
