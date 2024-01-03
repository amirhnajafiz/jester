package cmd

import "github.com/spf13/cobra"

type Agent struct{}

func (a Agent) Command() *cobra.Command {
	return nil
}

func (a Agent) main() {

}
