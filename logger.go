package main

import (
	"fmt"
	"time"
)

func LogAsync(message string) {

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("LOG:", message)
	}()
}