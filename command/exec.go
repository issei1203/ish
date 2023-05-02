package command

import "fmt"

func Exec(cmd string) {
	var command command

	err := parse(&cmd, &command)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s  ::  %s\n", command.name, command.option)
}
