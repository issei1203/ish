package command

import (
	"fmt"
	"ish/file"
)

type write struct{}

func (write write) exec(cmd parsedCommand) {
	if len(cmd.option) == 0 {
		fmt.Println("file name is not find")
		return
	} else if len(cmd.option) == 1 {
		fmt.Println("file name is not find")
		return
	}

	var fileName = cmd.option[len(cmd.option)-1]
	var content = cmd.option[len(cmd.option)-2]

	file.HEAD.WriteFile(fileName, content)
}
