package main

import (
	"code.google.com/p/go.crypto/ssh"
)

func main() {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)

}
