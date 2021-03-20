// main.go

package main


import (
	"fmt"
	"os"
	"os/user"
	"monkeyc/repl"
)

func main() {
	user, error := user.Current()
	if error != nil {
		panic(error)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")

	repl.Start(os.Stdin, os.Stdout)
}