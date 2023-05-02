package command

import (
	"fmt"
	"strings"
)

const space = " "

func parse(cmd *string, result *command) error {
	var splintedCmd = strings.Split(*cmd, space)
	var cmdLength = len(splintedCmd)

	if cmdLength == 0 {
		return fmt.Errorf("command is not exist")
	} else if cmdLength == 1 {
		*result = command{name: splintedCmd[0], option: []string{}}
	} else if cmdLength == 2 {
		*result = command{name: splintedCmd[0], option: []string{splintedCmd[1]}}
	} else {
		*result = command{name: splintedCmd[0], option: splintedCmd[1:cmdLength]}
	}

	return nil
}
