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
	fmt.Printf("Read config : %+v\n", cfg)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filePath := homeDir + "/.gatorconfig.json"

	cfg.SetUser("Misaque Ahmed", filePath)

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading JSON file: %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
