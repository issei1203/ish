package command

import "fmt"

var cmdMap = map[string]command{
	"ls": ls{},
}

func Exec(cmd string) {
	var command parsedCommand

	err := parse(&cmd, &command)
	if err != nil {
		fmt.Println(err)
		return
	}

	if executer, ok := cmdMap[command.name]; ok {
		executer.exec(command)
	}
}
