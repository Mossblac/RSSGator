package ext

type Commands struct {
	Handlers map[string]func(*State, Command) error
}
