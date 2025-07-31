package ext

func (c *Commands) Register(name string, handler func(*State, Command) error) {
	c.Handlers[name] = handler
}
