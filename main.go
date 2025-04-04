package main

import (
	"fmt"
	"log"
	"os"

	"github.com/binaryCoder-101/blog-aggregator/internal/config"
	_ "github.com/lib/pq"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("error: too few arguments")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command{
		commandName: cmdName,
		commandArgs: cmdArgs,
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading JSON file: %v", err)
	}
	// fmt.Printf("Read config : %+v\n", cfg)

	s := state{}

	s.configState = &cfg

	c := commands{
		commandHandler: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)

	err = c.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
