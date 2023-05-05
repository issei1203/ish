package command

import (
	"fmt"
	"ish/file"
)

type read struct{}

func (read read) exec(cmd parsedCommand) {
	if len(cmd.option) == 0 {
		fmt.Println("file name is not find")
		return
	}

	var fileName = cmd.option[len(cmd.option)-1]

	file.HEAD.ReadFile(fileName)
}
