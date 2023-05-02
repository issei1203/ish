package command

import (
	"fmt"
	"ish/file"
)

type mkdir struct{}

func (mkdir mkdir) exec(cmd parsedCommand) {
	if len(cmd.option) == 0 {
		fmt.Println("directory name is not exist")
		return
	}

	var dirName = cmd.option[len(cmd.option)-1]
	file.HEAD.MakeDir(dirName)
}
