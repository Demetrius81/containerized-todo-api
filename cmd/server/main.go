package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	// "github.com/joho/godotenv"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 		os.Exit(1)
// 	}
// }

func main() {
	connStr := getConnString()
	host := getHost()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		server, err := InitializeServer(connStr)
		if err != nil {
			log.Printf("can not initialize server: %v", err)
			os.Exit(1)
		}
		server.Start(host)
	}()

	<-ch
	fmt.Println("Shutting down...")
}

func getConnString() string {
	connStr, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Println("DATABASE_URL not found in environment variables")
		os.Exit(1)
	}

	return connStr
}

func getHost() string {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		log.Println("HOST not found in environment variables")
		os.Exit(1)
	}

	return ":" + host
}
