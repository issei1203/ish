package command

import (
	"fmt"
	"ish/file"
)

type rmdir struct{}

func (rmdir rmdir) exec(cmd parsedCommand) {
	if len(cmd.option) == 0 {
		fmt.Println("directory name is not find")
		return
	}

	var dirName = cmd.option[len(cmd.option)-1]
	file.HEAD.RemoveDir(dirName)
}
