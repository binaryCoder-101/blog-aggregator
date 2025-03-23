package main

import (
	"fmt"
	"log"
	"os"

	"github.com/binaryCoder-101/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading JSON file: %v", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filePath := homeDir + "/.gatorconfig.json"

	cfg.SetUser("Misaque Ahmed", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error reading json data: %v", err)
	}

	fmt.Printf("%+v", data)
}
