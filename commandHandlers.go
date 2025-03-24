package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.commandArgs) < 1 {
		return fmt.Errorf("error: username is required")
	}
	if len(cmd.commandArgs) > 1 {
		return fmt.Errorf("error: too many arguments")
	}

	if s.configState == nil {
		return fmt.Errorf("error: config is uninitialized")
	}
	s.configState.CurrentUserName = cmd.commandArgs[0]

	fmt.Printf("success: user has been set")

	return nil
}
