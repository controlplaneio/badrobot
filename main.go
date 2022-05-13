package main

import (
	"os"

	"github.com/controlplaneio/badrobot/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
