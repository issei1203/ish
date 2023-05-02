package command

import (
	"fmt"
	"ish/file"
)

type ls struct{}

func (ls ls) exec(cmd parsedCommand) {
	var contents = file.HEAD.GetContent()

	for _, content := range contents {
		var directoryFlag = "-"
		if content.IsDirectory {
			directoryFlag = "d"
		}
		fmt.Printf("%s %s\n", directoryFlag, content.Name)
	}
}
