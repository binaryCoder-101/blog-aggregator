package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/binaryCoder-101/blog-aggregator/internal/config"
	"github.com/binaryCoder-101/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

const dbURL = "postgres://postgres:postgresql@localhost:5432/gator"

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

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening connection to database: %v", err)
	}

	dbQueries := database.New(db)

	s := state{}

	s.configState = &cfg
	s.db = dbQueries

	c := commands{
		commandHandler: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)

	err = c.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
