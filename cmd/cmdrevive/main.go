package main

import (
	"github.com/narita-takeru/cmdrevive"
	"os"
	"strings"
)

func main() {

	targetDir := os.Args[1]
	pattern := os.Args[2]
	cmd := os.Args[3]
	args := os.Args[4:]

	dirs := strings.Split(targetDir, " ")

	cmdrevive.Start(dirs, pattern, cmd, args)
}
