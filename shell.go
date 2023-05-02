package main

import (
	"bufio"
	"fmt"
	"ish/command"
	"ish/file"
	"os"
)

func main() {
	initShell()

	renderLoop()

	exitShell()
}

func initShell() {
	file.InitHead()
}

func renderLoop() {
	fmt.Printf("%s > ", file.HEAD.GetCurrentDirName())

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	var cmd = stdin.Text()

	command.Exec(cmd)

	renderLoop()
}

func exitShell() {

}
