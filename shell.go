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
	fmt.Scanf("%s", &cmd)
	fmt.Println(cmd)

	renderLoop()
}

func exitShell() {

}
