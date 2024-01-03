package main

import (
	"log"

	"github.com/amirhnajafiz/jester/cmd"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{}

	root.AddCommand(
		cmd.Publisher{}.Command(),
		cmd.Consumer{}.Command(),
		cmd.Agent{}.Command(),
	)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
