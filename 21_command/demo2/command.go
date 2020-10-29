package main

import "fmt"

type Command interface {
	Execute()
}

type PingCommand struct{}

func (p *PingCommand) Execute() {
	fmt.Println("ping")
}

type StatusCommand struct{}

func (s *StatusCommand) Execute() {
	fmt.Println("status")
}

func execByName(name string) {
	command := map[string]Command{
		"ping":   &PingCommand{},
		"status": &StatusCommand{},
	}

	if c := command[name]; c != nil {
		c.Execute()
	} else {
		fmt.Println("No such command found")
	}
}

func main() {
	execByName("ping")
	execByName("status")
	execByName("test")
}
