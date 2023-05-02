package command

import "fmt"

var cmdMap = map[string]command{
	"ls":    ls{},
	"mkdir": mkdir{},
}

func Exec(cmd string) {
	var command parsedCommand

	err := parse(&cmd, &command)
	if err != nil {
		fmt.Println(err)
		return
	}

	if executor, ok := cmdMap[command.name]; ok {
		executor.exec(command)
	} else {
		fmt.Println("command not found")
	}
}
