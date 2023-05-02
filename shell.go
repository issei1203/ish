package main

import "fmt"

func main() {
	initShell()

	renderLoop()

	exitShell()
}

func initShell() {

}

func renderLoop() {
	fmt.Print("> ")

	var cmd string
	var _, err = fmt.Scanf("%s", &cmd)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cmd)

	renderLoop()
}

func exitShell() {

}
