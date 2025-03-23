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
		log.Fatal(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filePath := homeDir + "/.gatorconfig.json"

	cfg.SetUser("Misaque Ahmed", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("error reading json data")
	}

	fmt.Printf(string(data))
}
