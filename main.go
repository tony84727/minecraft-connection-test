package main

import (
	"fmt"
	"github.com/Craftserve/mcstatus"
	"log"
	"os"
	"time"
)

func printUsage() {
	fmt.Println("minecraft-connection-test <address>")
}

func main() {
	if len(os.Args) <= 1 {
		printUsage()
	}
	address := os.Args[1]
	addr, err := mcstatus.Resolve(address)
	if err != nil {
		log.Fatalf("invalid address: %s\n", address)
	}
	for i := 0; i < 25; i++ {
		_, latency, err := mcstatus.CheckStatus(addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("latency: %v\n", latency)
		time.Sleep(5 * time.Second)
	}
}
