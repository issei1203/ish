package command

type parsedCommand struct {
	name   string
	option []string
}

type command interface {
	exec(cmd parsedCommand)
}
