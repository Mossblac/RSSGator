package ext

import "fmt"

func (c *Commands) Run(s *State, cmd Command) error {
	if handler, ok := c.Handlers[cmd.CommandName]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.CommandName)
}
