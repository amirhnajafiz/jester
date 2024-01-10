package main

import (
	"github.com/amirhnajafiz/jester/cmd"
	"github.com/amirhnajafiz/jester/internal/config"

	"github.com/spf13/cobra"
)

func main() {
	// root command
	root := &cobra.Command{}

	// load configs
	cfg := config.Load("config.yml")

	root.AddCommand(
		cmd.Publisher{}.Command(),
		cmd.Consumer{}.Command(),
		cmd.Agent{
			Cfg: cfg,
		}.Command(),
	)

	// execute root command
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
