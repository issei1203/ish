package command

import (
	"fmt"
	"ish/file"
)

type rm struct{}

func (rm rm) exec(cmd parsedCommand) {
	if len(cmd.option) == 0 {
		fmt.Println("file name is not find")
		return
	}

	var fileName = cmd.option[len(cmd.option)-1]
	file.HEAD.RemoveFile(fileName)
}
