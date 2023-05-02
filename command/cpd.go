package command

import "ish/file"

type cpd struct{}

func (cpd cpd) exec(cmd parsedCommand) {
	file.HEAD.ChangeParentDir()
}
