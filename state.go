package main

import (
	"github.com/binaryCoder-101/blog-aggregator/internal/config"
	"github.com/binaryCoder-101/blog-aggregator/internal/database"
)

type state struct {
	db          *database.Queries
	configState *config.Config
}
