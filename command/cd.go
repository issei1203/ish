package command

import (
	"fmt"
	"ish/file"
)

type cd struct{}

func (cd cd) exec(cmd parsedCommand) {
	if len(cmd.option) == 0 {
		fmt.Println("No such file or directory")
		return
	}

	var dirName = cmd.option[len(cmd.option)-1]
	file.HEAD.ChangeDir(dirName)
}
