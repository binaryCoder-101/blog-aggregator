package main

import (
	"fmt"
	"strings"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.commandArgs) == 0 {
		return fmt.Errorf("error: username is required")
	}

	if s.configState == nil {
		return fmt.Errorf("error: config is uninitialized")
	}

	username := strings.Join(cmd.commandArgs[0:], " ")

	err := s.configState.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("success: user has been set")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	return nil
}
