package main

import (
	"github.com/narita-takeru/cmdrevive"
	"os"
)

func main() {

	targetDir := os.Args[1]
	pattern := os.Args[2]
	cmd := os.Args[3]
	args := os.Args[4:]

	cmdrevive.Start(targetDir, pattern, cmd, args)
}
